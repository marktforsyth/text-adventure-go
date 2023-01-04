package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type player struct {
	Health int
	CurrentRoom string
	GoldPieces int
	Inventory []storeItem
}

func main() {
	rooms := fetchRooms()
	store := fetchStore()

	mainPlayer := player{
		Health: 70,
		CurrentRoom: rooms[0].Name,
		GoldPieces: 0,
		Inventory: []storeItem{},
	}

	fmt.Println("\nHello! Welcome to this text adventure. You are in the " + mainPlayer.CurrentRoom)
	fmt.Println(rooms[indexOfRoom(rooms, mainPlayer.CurrentRoom)].Description)
	fmt.Println()
	printInstructions()

	runGame(mainPlayer, rooms, store)
}

func runGame(mainPlayer player, rooms []room, store []storeItem) {
	editPlayer := &mainPlayer

	// Display current stats
	fmt.Println("Health: " + strconv.Itoa(mainPlayer.Health))
	fmt.Println("Gold Pieces: " + strconv.Itoa(mainPlayer.GoldPieces))
	fmt.Println()

	// Display options in the current room
	fmt.Println("Here are your options in the " + mainPlayer.CurrentRoom)
	for n, option := range rooms[indexOfRoom(rooms, mainPlayer.CurrentRoom)].Options {
		if option.AlreadyDone {
			fmt.Println("    " + strconv.Itoa(n + 1) + ") " + option.Description + " [Already Done]")
		} else {
			fmt.Println("    " + strconv.Itoa(n + 1) + ") " + option.Description)
		}
	}

	// Get user input
	scanner := bufio.NewScanner(os.Stdin)
	var userInput string
	if scanner.Scan() {
		userInput = scanner.Text()
	}

	checkUserInput(userInput, rooms, store, mainPlayer, editPlayer)

	if userInput != "exit" && mainPlayer.Health > 0 {
		runGame(mainPlayer, rooms, store)
	} else {
		return
	}
}

func indexOfRoom(rooms []room, roomName string) int {
	for i, room := range rooms {
		if (room.Name == roomName) {
			return i
		}
	}
	return -1
}

func checkUserInput(userInput string, rooms []room, store []storeItem, mainPlayer player, editPlayer *player) {
	fmt.Println()

	if userInput == "1" || userInput == "2" || userInput == "3" {
		userChoice, err := strconv.Atoi(userInput)
		if err != nil {
			fmt.Println(err)
		}

		if !rooms[indexOfRoom(rooms, mainPlayer.CurrentRoom)].Options[userChoice - 1].AlreadyDone {
			fmt.Println(rooms[indexOfRoom(rooms, mainPlayer.CurrentRoom)].Options[userChoice - 1].Result)
			if rooms[indexOfRoom(rooms, mainPlayer.CurrentRoom)].Options[userChoice - 1].ImportantResult != "" {
				reactToImportantResults(editPlayer, rooms[indexOfRoom(rooms, mainPlayer.CurrentRoom)].Options[userChoice - 1].ImportantResult)
			}

			rooms[indexOfRoom(rooms, mainPlayer.CurrentRoom)].Options[userChoice - 1].AlreadyDone = true
		} else {
			fmt.Println("You already did this, rememeber? I like to be petty, so I won't allow you to do it again.")
		}
	} else if userInput == "list rooms" {
		fmt.Println("Adjacent Rooms:")
		for _, room := range rooms[indexOfRoom(rooms, mainPlayer.CurrentRoom)].AdjacentRooms {
			fmt.Println("    " + room)
		}
	} else if strings.HasPrefix(userInput, "enter") {
		enterRoom(userInput[6:len(userInput)], rooms, editPlayer)
	} else if userInput == "store" {
		numOfAlreadyBoughtItems := 0

		fmt.Println("Items in the Store:")
		for _, storeItem := range store {
			if storeItem.AlreadyBought {
				numOfAlreadyBoughtItems ++
				fmt.Println("    " + storeItem.Name + " (" + strconv.Itoa(storeItem.Price) + " GP) - " + storeItem.Description + " [Already Bought]")
			} else {
				fmt.Println("    " + storeItem.Name + " (" + strconv.Itoa(storeItem.Price) + " GP) - " + storeItem.Description)
			}
		}

		if numOfAlreadyBoughtItems == len(store) {
			fmt.Println("(It looks like you've already bought everything there is to buy in the store.)")
		}
	} else if strings.HasPrefix(userInput, "buy") {
		buyFromStore(userInput[4:len(userInput)], store, editPlayer)
	} else if userInput == "inventory" {
		if len(mainPlayer.Inventory) == 0 {
			fmt.Println("There is nothing in your inventory.")
		} else {
			fmt.Println("Items in Inventory:")
			for _, inventoryItem := range mainPlayer.Inventory {
				fmt.Println("    " + inventoryItem.Name + " - " + inventoryItem.Description)
			}
		}
	} else if userInput == "help" {
		printInstructions()
	} else if userInput == "exit" {
		return
	} else {
		if userInput != "" {
			fmt.Println("Error: command not found. Did you misspell something?")
		}
	}

	if rooms[indexOfRoom(rooms, mainPlayer.CurrentRoom)].Options[0].AlreadyDone &&
	rooms[indexOfRoom(rooms, mainPlayer.CurrentRoom)].Options[1].AlreadyDone &&
	rooms[indexOfRoom(rooms, mainPlayer.CurrentRoom)].Options[2].AlreadyDone {
		fmt.Println("(It looks like you've already done everything there is to do in this room. Why not move onto another one?)")
	}

	fmt.Println()
}

func printInstructions() {
	fmt.Println("Commands:")
	fmt.Println("1 - select option one in a room")
	fmt.Println("2 - select option two in a room")
	fmt.Println("3 - select option three in a room")
	fmt.Println("list rooms - display all of the adjacent rooms")
	fmt.Println("enter [room name] - enter an adjacent room with the name [room name]")
	fmt.Println("store - open up the unhelpful store to see what useless accessories you can buy")
	fmt.Println("buy [item name] - buy a useless accessory --you can do this at any time if you have the GP")
	fmt.Println("inventory - open up your inventory to see the useless things you've bought")
	fmt.Println("exit - exit the game")
	fmt.Println("help - show these instructions again")
	fmt.Println()
}

func reactToImportantResults(editPlayer *player, result string) {
	if result == "gain money" {
		editPlayer.GoldPieces ++
	} else if result == "lose money" && editPlayer.GoldPieces > 0 {
		editPlayer.GoldPieces --
	} else if result == "gain health" {
		editPlayer.Health += 10

		if editPlayer.Health >= 100 {
			editPlayer.Health = 100
		}
	} else if result == "lose health" {
		editPlayer.Health -= 10

		if editPlayer.Health <= 0 {
			fmt.Println("You died. Goodbye. It was nice knowing you!")
		}
	} else {
		fmt.Println("Error in reactToImportantResults(): the result doesn't match a category. Result: " + result)
	}
}

func enterRoom(roomName string, rooms []room, editPlayer *player) {
	canEnterRoom := false

	for _, room := range rooms {
		if room.Name == roomName {
			for _, adjacentRoom := range rooms[indexOfRoom(rooms, editPlayer.CurrentRoom)].AdjacentRooms {
				if room.Name == adjacentRoom {
					canEnterRoom = true
				}
			}
		}
	}

	if canEnterRoom {
		editPlayer.CurrentRoom = roomName
		fmt.Println("You've entered the " + roomName + "!")
		fmt.Println(rooms[indexOfRoom(rooms, editPlayer.CurrentRoom)].Description)
	} else {
		fmt.Println("Error: '" + roomName + "' is not an adjacent room.")
	}
}

func buyFromStore(itemName string, store []storeItem, editPlayer *player) {
	itemIsInStore := false
	var itemIndex int

	for i, storeItem := range store {
		if storeItem.Name == itemName {
			itemIsInStore = true
			itemIndex = i
		}
	}

	if itemIsInStore {
		if store[itemIndex].AlreadyBought == true {
			fmt.Println("You already bought that, and there was only one, so you can't buy it again.")
			return
		} else if editPlayer.GoldPieces < store[itemIndex].Price {
			fmt.Println("Error: you do not have enough money to buy that.")
			return
		}

		editPlayer.GoldPieces -= store[itemIndex].Price
		editPlayer.Inventory = append(editPlayer.Inventory, store[itemIndex])
		store[itemIndex].AlreadyBought = true
		fmt.Println("You have successfully bought the " + store[itemIndex].Name + "!")
	} else {
		fmt.Println("Error: that item is not in the store, so you cannot buy it.")
	}
}