package services

import (
	"TEST/quanngo/hackathon/storage/postgres"
	"TEST/quanngo/hackathon/types"

	"fmt"
	"image"
	"image/jpeg"
	"mime/multipart"
	"os"

	"xorm.io/xorm"
)

type User struct {
	UserName string
	Password string
}

type File struct {
	FileName string
	Type     string
	Size     int64
}

func CreateUser(args *types.Register) error {
	return postgres.InTransaction(func(session *xorm.Session) error {
		User := User{
			UserName: args.UserName,
			Password: args.Password,
		}

		if _, err := session.Insert(&User); err != nil {
			return err
		}

		return nil
	})
}

func CheckUser(email string) error {
	rawSQL := `
	SELECT *
	FROM ` + postgres.Dialect.Quote("user") + `
	WHERE user_name = ? `

	var user types.Register
	exists, err := postgres.X.SQL(rawSQL, email).Get(&user)
	if err != nil {
		return err
	}
	if !exists {
		User := types.Register{
			UserName: email,
			Password: "123456",
		}
		err := CreateUser(&User)
		if err != nil {
			return err
		}

		return nil
	}

	return nil
}

func CreateFile(fileUpload multipart.File, args *types.FileInfo) error {

	file, _, err := image.Decode(fileUpload)
	if err != nil {
		return err
	}

	dirName := "images/"
	if _, err := os.Stat(dirName); err != nil {
		err = os.MkdirAll(dirName, os.ModePerm)
		if err != nil {
			return err
		}
	}

	path := fmt.Sprintf("%s%s", dirName, args.FileName)
	dst, err := os.Create(path)
	if err != nil {
		return err
	}
	defer dst.Close()

	var opt jpeg.Options
	opt.Quality = 50
	err = jpeg.Encode(dst, file, &opt)
	if err != nil {
		return err
	}

	var sizeLimit int64 = 8 * 1024 * 1024 //MB
	fileStat, err := os.Stat(path)
	if err != nil {
		return err
	}
	if fileStat.Size() > sizeLimit {
		dst.Close()
		if err := os.Remove(path); err != nil {
			return err
		}
		return fmt.Errorf("File Is Over Limit")
	}

	err = SaveFileInfo(args)
	if err != nil {
		return err
	}

	return nil
}

func SaveFileInfo(args *types.FileInfo) error {
	return postgres.InTransaction(func(session *xorm.Session) error {
		File := File{
			FileName: args.FileName,
			Size:     args.Size,
			Type:     args.Type,
		}

		if _, err := session.Insert(&File); err != nil {
			return err
		}

		return nil
	})
}
