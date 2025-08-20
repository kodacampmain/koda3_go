package minitask1

import "errors"

type users struct {
	email    string
	password string
}

var userList = [2]users{
	{email: "a@mail.com", password: "passa"},
	{email: "b@mail.com", password: "passb"},
}

func Login(email, password string) (string, error) {
	for _, v := range userList {
		if v.email == email {
			if v.password != password {
				continue
			}
			return "Login Berhasil", nil
		}
	}
	return "", errors.New("login gagal")
}
