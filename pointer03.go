// http://openmymind.net/Things-I-Wish-Someone-Had-Told-Me-About-Go/

package main

type User struct {
  Name string
}

func Modify1(u *User) {
  u = &User{Name: "Paul"}
}

func Modify2(u User) {
  u.Name = "Duncan"
}

func Modify3(u **User) {
 *u = &User{Name: "Bob"}
}

func main() {
  u1 := &User{Name: "Leto"}
  println(u1.Name)    // Leto
  Modify1(u1)
  println(u1.Name)    // Leto

  u2 := User{Name: "Leto"}
  println(u2.Name)    // Leto
  Modify2(u2)
  println(u2.Name)    // Leto

  u3 := &User{Name: "Leto"}
  println(u3.Name)    // Leto
  Modify3(&u3)
  println(u3.Name)    // Bob
}
