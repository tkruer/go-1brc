package main

import ( 
    "encoding/csv"
    "fmt"
    "log"
    "os"
) 

func main() {
  // Let's get started and open the file /data/weather_station.csv
  
   file, err := os.Open("data/weather_stations.csv") 
      
    if err != nil { 
        log.Fatal("Error while reading the file", err) 
    } 
  
    defer file.Close() 
  
    reader := csv.NewReader(file) 
      
    records, err := reader.ReadAll() 
  
    if err != nil { 
        fmt.Println("Error reading records") 
    } 
      
    for _, eachrecord := range records { 
        fmt.Println(eachrecord) 
    } 
}
