package purchase

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
    return kind == "car" || kind == "truck"
	panic("NeedsLicense not implemented")
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
    option := option1
    if option1 > option2 {
        option = option2
    }
    return option + " is clearly the better choice."
	panic("ChooseVehicle not implemented")
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
    ageInt := int(age)
    if ageInt < 3 {
        return (originalPrice * float64(80))/float64(100)
    }else if ageInt >= 3 && ageInt < 10{
         return (originalPrice * float64(70))/float64(100)
    }else {
        return (originalPrice * float64(50))/float64(100)
    }
	panic("CalculateResellPrice not implemented")
}
