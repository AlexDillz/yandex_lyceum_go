// package main

// import (
// 	"fmt"
// )

// func main() {

// 	for i := 0; i <= 10; i += 2 {
// 		fmt.Println(i)
// 	}

// }

// цикл for в режиме for

// package main

// import (
// 	"fmt"
// )

// func main() {

// for i := 0; i < 10; i++ {
//     for j := 0; j < 10; j++ {
//         if i * j > 20 {
//             continue
//         }
//         fmt.Println(i * j)
//     }
// }
// }

//ОБРАТИТЬ ВНИМАНИЕ
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	for i := 0; i < 10; i++ {
// 		for j := 0; j < 10; j++ {
// 			if i*j > 20 {
// 				continue
// 			}
// 			fmt.Println(i * j)
// 		}
// 	}
// }

package main

import (
	"fmt"
)

func main() {
	for i, letter := range "Hello, world!" {
		fmt.Println(i, string(letter))
	}
}

//Убрали i в выводе, в инициализации заменили i на _
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	for _, letter := range "Hello, world!" {
// 		fmt.Println(string(letter))
// 	}
// }
