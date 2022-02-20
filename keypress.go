package main

import "os/exec"

type keytool struct{}

func (u *keytool) run(c ...string) []byte {
	out, _ := exec.Command("/usr/bin/xdotool", c...).Output()
	return out
}
func (u *keytool) key(val string, n int) {
	if n != 0 {
		for i := 0; i < n; i++ {
			u.run(val)
		}
	}
}
func (u *keytool) Up(n int) *keytool {
	u.key("Up", n)
	return u
}
func (u *keytool) Down(n int) *keytool {
	u.key("Down", n)
	return u
}
func (u *keytool) Right(n int) *keytool {
	u.key("Right", n)
	return u
}
func (u *keytool) Left(n int) *keytool {
	u.key("Left", n)
	return u
}
func (u *keytool) Return(n int) *keytool {
	u.key("Return", n)
	return u
}
