package main

func ColorToHex(color string) string {
	var hex string
	switch color {
	case "green":
		hex = "32c12c"
	case "teal":
		hex = "009888"
	case "indigo":
		hex = "3e49bb"
	case "blue":
		hex = "526eff"
	case "purple":
		hex = "7f4fc9"
	case "lightgreen":
		hex = "87c735"
	case "lime":
		hex = "cde000"
	case "lightblue":
		hex = "00a5f9"
	case "cyan":
		hex = "00bcd9"
	case "darkpurple":
		hex = "682cbf"
	case "yellow":
		hex = "ffef00"
	case "orange":
		hex = "ff9a00"
	case "lightred":
		hex = "ff9a00"
	case "brown":
		hex = "7c5547"
	case "bluegrey":
		hex = "5f7d8e"
	case "amber":
		hex = "ffcd00"
	case "darkorange":
		hex = "ff5500"
	case "red":
		hex = "d40c00"
	case "darkbrown":
		hex = "50342c"
	case "grey":
		hex = "9e9e9e"
	case "white":
		hex = "ffffff"
	case "black":
		hex = "000000"
	default:
		hex = color
	}
	return hex
}
