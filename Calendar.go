package main

import (
    "fmt"
    "time"
)

func validateMonthlyRepeat(startDate, endDate time.Time, listOfDays map[int]bool ) bool{
    if(startDate.Equal(endDate)){
        day := startDate.Day()
        _,available := listOfDays[day]
        if(available){
            fmt.Println("log: validateMonthlyRepeat() available-",available)
            return true
        }
    }
    pointingDate := startDate
    for(endDate.After(pointingDate)){
        fmt.Println("log:pointingDate:", pointingDate)
        pointingDate = pointingDate.AddDate(0, 0, 1)
        day := pointingDate.Day()
        _,available := listOfDays[day]
        if(available){
            fmt.Println("log: validateMonthlyRepeat() available-",pointingDate)
            return true
        }
    }
    return false
}

func validTime(year, month, day, hour, minute, second, ns int) (bool, time.Time) {
    date1 := time.Date(year, time.Month(month), day, hour, minute, second, ns, time.UTC)
    fmt.Println("checkTime() log: ",date1)
    actualMonth := int(date1.Month())
    actualDay := int(date1.Day())
    if(actualMonth==month){
        return true, date1
    }
    if(actualDay==day){
        return true, date1
    }
    return false, date1
}

func main() {
    //Inputs: 
    validStartDate, startDate := validTime(2022, 2, 1, 10, 30, 0, 0)
    validEndDate, endDate := validTime(2022, 3, 28, 10, 30, 0, 0)
    listOfDays := map[int]bool{30: true, 31: true, 29: true}
    
    //LOGIC
    if(validStartDate && validEndDate){
        fmt.Println("log: both start date & end date are valid: [", startDate,"] , [", endDate,"]")
        result := validateMonthlyRepeat(startDate, endDate, listOfDays)
        fmt.Println("RESULT: validateMonthlyRepeat()=",result)
    } else if(validStartDate==false && validEndDate==false){
        fmt.Println("ERROR: Invalid start Date and end Date!")
    } else if(validStartDate==false){
        fmt.Println("ERROR: Invalid start Date!")
    } else{
        fmt.Println("ERROR: Invalid end Date!")
    }
    
}
//https://www.mycompiler.io/new/go

