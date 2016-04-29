package main

type User struct {
    Name string
}

func main() {
    u := &User{Name: "Leto"}
    println(u.Name)

    Modify(u)
    println(u.Name)

    Modify1(u)
    println(u.Name)

    Modify2(&u)
    println(u.Name)    
}

func Modify2(u **User) {
    (*u).Name = "Jim"
    // *u = &User{Name: "Bob"}
}

func Modify1(u *User) {
    u.Name = "Paul"
}

func Modify(u *User) {
    u = &User{Name: "Paul"}
}
