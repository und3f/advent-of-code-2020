package main

import (
	"fmt"
	//"io/ioutil"
	//"strconv"
	//"strings"
)

const subjectNumber = 5764801

func transformSubjectNumber(value uint64, subjectNumber uint64) uint64 {
	value = value * subjectNumber
	value = value % 20201227

	return value
}

func transformNTimes(value, subjectNumber, times uint64) uint64 {
	for i := uint64(0); i < times; i++ {
		value = transformSubjectNumber(value, subjectNumber)
	}

	return value
}

const startValue = uint64(1)
const publicKeySubjectNumber = uint64(7)

func handShake(cardSecretLoopSize, doorSecretLoopSize uint64) bool {
	cardPublicKey := transformNTimes(startValue, publicKeySubjectNumber, cardSecretLoopSize)
	doorPublicKey := transformNTimes(startValue, publicKeySubjectNumber, doorSecretLoopSize)

	fmt.Println("handShake", cardSecretLoopSize, doorSecretLoopSize)
	fmt.Println("	public key", cardPublicKey, doorPublicKey)

	encryptionKey := transformNTimes(1, doorPublicKey, cardSecretLoopSize)
	encryptionKey2 := transformNTimes(1, cardPublicKey, doorSecretLoopSize)
	fmt.Println("encryption1:", encryptionKey)
	fmt.Println("encryption2:", encryptionKey2)

	return false
}

func determineDoorLoopSize(expectedValue uint64) uint64 {
	value := startValue
	var i uint64

	for value != expectedValue {
		i++
		value = transformSubjectNumber(value, publicKeySubjectNumber)
	}

	return i
}

const cardPublicKey = uint64(1965712)
const doorPublicKey = uint64(19072108)

func main() {
	var partTwo uint64

	doorLoopSize := determineDoorLoopSize(doorPublicKey)
	cardLoopSize := determineDoorLoopSize(cardPublicKey)
	fmt.Println("card loop size", cardLoopSize)
	fmt.Println("door loop size", doorLoopSize)

	encryptionKey := transformSubjectNumber(cardLoopSize, doorPublicKey)
	fmt.Println("Part One:", encryptionKey)
	encryptionKey = transformSubjectNumber(doorLoopSize, cardPublicKey)
	fmt.Println("Part One:", encryptionKey)

	fmt.Println("Handshake", handShake(cardLoopSize, doorLoopSize))

	fmt.Println("Part Two:", partTwo)
}
