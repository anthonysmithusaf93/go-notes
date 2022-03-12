package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {
  //Generate random seed with time in unix nano seconds
  rand.Seed(time.Now().UnixNano())
  //Declare heist on
  isHeistOn := true
  //Random int to determine if we made it past the guards
  eludedGuards := rand.Intn(100)
  //Sneak past guards conditional
  if eludedGuards >= 50 {
    fmt.Println("Looks like you've managed to make it past the guards. Good job, but remember, this is the first step.")
    } else {
      fmt.Println("Plan a better disguise next time?")
    }
  //Open vault conditional
  openedVault := rand.Intn(100)
  if openedVault >= 70 {
    fmt.Println("Grab and GO!")
  } else if openedVault < 70 {
    isHeistOn = false
    fmt.Println("We couldn't open the vault, let's get out of here!")
  }
  //Escape conditional
  if isHeistOn {
    switch leftSafely := rand.Intn(5); leftSafely {
      case 0:
        isHeistOn = false
        fmt.Println("We failed the heist!")
      case 1:
        isHeistOn = false
        fmt.Println("Really, a robo guard?")
      case 2:
        isHeistOn = false
        fmt.Println("How did you trip over the one alarm wire...?")
      case 3:
        isHeistOn = false
        fmt.Println("Who would have thought we couldn't open the vault from the inside...")
      default:
        fmt.Println("Start the getaway car!")
    }
    if isHeistOn {
      amtStolen := 10000 + rand.Intn(1000000)
      fmt.Printf("We stole $%d!!\n", amtStolen)
    }
  }
  fmt.Println(isHeistOn)
}
