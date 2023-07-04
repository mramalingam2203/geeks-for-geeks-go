// https://www.geeksforgeeks.org/mathematical-algorithms/?ref=shm

package main()

import "fmt"

func main() {

}

func gcd(a, b int){
    if b == 0{
        return a
	}
    return gcd(b, a % b)
}