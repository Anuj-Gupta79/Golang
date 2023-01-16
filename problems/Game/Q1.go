package main

// import "fmt"

func CanFastAttack(knightIsAwake bool) bool {
	return !knightIsAwake
}

func CanSpy(knightIsAwake, archerIsAwake, prisonerIsAwake bool) bool {
	return knightIsAwake || archerIsAwake || prisonerIsAwake
}

func CanSignalPrisoner(archerIsAwake, prisonerIsAwake bool) bool {
	return !archerIsAwake && prisonerIsAwake
}

func CanFreePrisoner(knightIsAwake, archerIsAwake, prisonerIsAwake, petDogIsPresent bool) bool {
	if petDogIsPresent {
		return !knightIsAwake
	}
	if !knightIsAwake && !archerIsAwake {
		return prisonerIsAwake
	}
	return false
}

func main(){
	// var knightIsAwake bool
	// var archerIsAwake bool
	// var prisonerIsAwake bool
	// var petDogIsAwake bool
	
	// fmt.Scanf("%d %d %d %d", &knightIsAwake, &archerIsAwake, &prisonerIsAwake, &petDogIsAwake)

	// fmt.Println(CanFastAttack(knightIsAwake))


}
