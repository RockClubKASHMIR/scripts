/***** This script is created by RockClubKASHMIR <discord @RockClubKASHMIR#8058> *****\
 
 v5.0
 
    DESCRIPTION
 1. Always Keeps reseirved slot
 2. You can send EXPO fleets from more than 1 planet/moon 
 3. Check for EXPO Debris and recycle them (if you are Discoverer and have Pathfinders)
 4. You can set up your ship list by 2 methods (or by combination of both of them):
    a. All ships with quantity 0 that you set will be calculated automatically (full quantity divided by the free EXPO slots)
       - if sendAtOnce = true, all ships set with quantity 0 will be sent at once.
    b. Ships set with quantity different than 0 that you set will be accepted literally, and if any of your ships is even 1 less, the fleet will not be sent.
 5. You can start the script at certain time. Sending of the fleets will stop after repeats that you set.
 6. Evenly distribution of EXPO slots per each moon/planet (can be turn on/of) 
 */

homes = ["M:1:2:3"] // Replace M:1:2:3 with your coordinate - M for the moon, P for planet.
// You can add as many planets/moons you want - the home list must look like this: homes = ["M:1:2:3", "M:2:2:3"]

shipsList = {LARGECARGO: 0, LIGHTFIGHTER: 0, PATHFINDER: 1}// Set your Ships list

splitSlots = true //Do you want evenly distribution of EXPO slots per each moon/planet? true = YES / false = NO
sendAtOnce = false //Do you want to send the ships set with quantity 0 at once? true = YES / false = NO

minusCurrentSystem = 5 // Set this as start destination of range coordinates - minus your current world's system
plusCurrentSystem = 5 // Set this as end destination of range coordinates - plus your current world's system

DurationOfExpedition = 1 // Set duration (in hours) of the EXPEDITION: minimum 1 - maximum 8
PathfindersDebris = true // Do you want to get EXPO debrises? true = YES / false = NO
Pnbr = 5  // The script will ignore debris less than for PATHFINDERS that you set - The Maximum PATHFINDERS is limited only of your PATHFINDERS on the current moon/planet! You can set this value from 1, to the number you want
PathfinderSystemsRange = true // Do you want to check/get EXPO debris in range systems? true = YES / false = NO
SystemsRange = false // Do you want to send your EXPO fleet to Range coordinates? true = YES / false = NO
Repeat = true // Do you want to repeat the full cycle of fleet sending? true = YES / false = NO
HowManyCycles = 5 // Set the limit of repeats of whole cycle of EXPO fleet sending - 0 means forewer

myTime = "09:33:00"// Set your start Time; Hour: 00 - 23, Minute: 00 - 59
useStartTime = false // Do you want to run this script at specific time every day? true = YES / false = NO

//-------
current = 0
wrong = []
curentco = {}
homeworld = nil
i = 0
ei = 0
er = nil
err = nil
flag = 0
cng = 0
cycle = 0
endFlag = 0
fleetFlag = 0
RepeatTimes = 1
if (Pnbr < 1) {Pnbr = 1}
for home in homes {
    for celestial in GetCachedCelestials() {
        if GetCachedCelestial(celestial) == GetCachedCelestial(homes[0]) {
            homeworld = GetCachedCelestial(homes[i])
            ei = ei + 1
        } else {flag = 1}
    }
    if flag == 1 {wrong += homes[i]}
    if i < len(homes)-1 {
        i++
        flag = 0
    }
}
if ei == len(homes) {homeworld = GetCachedCelestial(homes[0])}
if !IsDiscoverer() {
    Print("You are not Discoverer and cannot get the EXPO Debris!")
    PathfindersDebris = false
}
if useStartTime == false {
    hour, minute, sec = Clock()
    startHour = hour
    startMin = minute
    startSec = sec + 3
    if startSec >= 60 {
        startSec = startSec - 60
        startMin = startMin + 1
        if startMin >= 60 {
            startMin = startMin - 60
            startHour = startHour + 1
        }
        if startHour >= 24 {startHour = startHour - 24}
    }
    myTime = ""+startSec+" "+startMin+" "+startHour+" * * *"
}
if HowManyCycles == 0 {HowManyCycles = false}
if homeworld != nil {
    CronExec(myTime, func() {
        slotMarker = 0
        totalUsl = GetSlots().Total - GetFleetSlotsReserved()
        totalExpSlots = GetSlots().ExpTotal
        Sleep(1500)
        for home = current; home <= len(homes)-1; home++ {
            pp = 0
            Dtarget = 0
            marker = home
            Expp = GetSlots().ExpInUse
            Sleep(800)
            homeworld = GetCachedCelestial(homes[home])
            if homeworld.Coordinate.IsMoon() {
                Print("Your Moon is: "+homeworld.Coordinate)
            } else {Print("Your Planet is: "+homeworld.Coordinate)}
            fromSystem = homeworld.GetCoordinate().System - minusCurrentSystem
            if fromSystem < 1 {fromSystem = 1}
            toSystem = homeworld.GetCoordinate().System + plusCurrentSystem
            if fromSystem > 499 {toSystem = 499}
            crdn = fromSystem
            ExpsTemp = 0
            if SystemsRange == true && cycle >= len(homes)-1 {
                for id, num in curentco {
                    if id == homes[home] {crdn = num}
                }
            }
            totalSlots = totalUsl
            slots = GetSlots().InUse
            Sleep(800)
            currentTime = 1
            
            if slots < totalSlots {
                slots = Expp
                totalSlots = totalExpSlots
                times = totalExpSlots - slots
                ExpsTemp = 1
                if slots == totalSlots {fleetFlag = 2}
            } else {
                fleetFlag = 1
                slots = totalSlots
            }
            Sleep(Random(1000, 3000))
            if err != nil {slots = totalSlots}
            if slots < totalSlots {
                times = totalExpSlots - Expp
                if splitSlots == true {
                    slotMarker = totalExpSlots-marker
                    times = slotMarker/len(homes)
                    if times > Floor(times) {times = Floor(times) + 1}
                    if times < 1 {times = 1}
                }
                if sendAtOnce == true {times = 1}
                Print(times+" slots will be used")
                if splitSlots == false && sendAtOnce == false {
                    times = totalExpSlots
                    currentTime = Expp
                }
                for time = currentTime; time <= times; time++ {
                    myShips, _ = homeworld.GetShips()
                    tt = 0
                    rtt = 0
                    ExpFleet = {}
                    if ExpsTemp == 0 {
                        totalSlots = totalUsl
                        slots = GetSlots().InUse
                        Sleep(800)
                        if slots < totalSlots {
                            slots = Expp
                            totalSlots = totalExpSlots
                            if slots == totalSlots {fleetFlag = 2}
                        } else {
                            fleetFlag = 1
                            slots = totalSlots
                        }
                    }
                    if err != nil {slots = totalSlots}
                    if slots < totalSlots {
                        ExpsTemp == 0
                        if SystemsRange == false {
                            Dtarget, _ = ParseCoord(homeworld.GetCoordinate().Galaxy+":"+homeworld.GetCoordinate().System+":"+16)
                        }
                        if SystemsRange == true {
                            if crdn > toSystem {crdn = fromSystem}
                            Dtarget, _ = ParseCoord(homeworld.GetCoordinate().Galaxy+":"+crdn+":"+16)
                        }
                        explist = []
                        Sleep(Random(13000, 18000)) // For avoiding ban
                        Flts, _ = GetFleets()
                        fleet = NewFleet()
                        fleet.SetOrigin(homeworld)
                        fleet.SetDestination(Dtarget)
                        fleet.SetSpeed(HUNDRED_PERCENT)
                        fleet.SetMission(EXPEDITION)
                        if len(shipsList) > 0 {
                            for ShipID, num in shipsList {
                                rtt = rtt + 1
                                fleetInAir = 0
                                if myShips.ByID(ShipID) != 0 {
                                    if num == 0 {
                                        if sendAtOnce == false {
                                            for f in Flts {
                                                ships = f.Ships
                                                if f.Mission == EXPEDITION {
                                                    if homeworld.Coordinate == f.Origin {
                                                        if ships.ByID(ShipID) != 0 {
                                                            fleetInAir = fleetInAir + ships.ByID(ShipID)
                                                        }
                                                    }
                                                }
                                            }
                                            fleetInAir = fleetInAir + myShips.ByID(ShipID)
                                            num = Floor(fleetInAir/times)
                                        }
                                        if sendAtOnce == true {num = myShips.ByID(ShipID)}
                                        if num > 0 {
                                            ExpFleet[ShipID] = num
                                            tt = tt + 1
                                        }
                                    } else {
                                        if ShipID != PATHFINDER {
                                            if myShips.ByID(ShipID) >= num {
                                                ExpFleet[ShipID] = num
                                                tt = tt + 1
                                            }
                                        } else {
                                            if myShips.ByID(ShipID) < num {num = myShips.ByID(ShipID)}
                                            ExpFleet[ShipID] = num
                                            tt = tt + 1
                                        }
                                    }
                                }
                            }
                        }
                        fleet.SetDuration(DurationOfExpedition)
                        if rtt == tt {
                            for ShipID, nbr in ExpFleet {
                                fleet.AddShips(ShipID, nbr)
                                explist += ShipID+": "+nbr
                            }
                        }
                        a, err = fleet.SendNow()
                        if err == nil {
                            cng = 1
                            slots = slots + 1
                            Expp = slots
                            if sendAtOnce == true {er = "no ships to send"}
                            Print(explist+" are sended successfully to "+Dtarget)
                            if SystemsRange == true {
                                if crdn <= toSystem {crdn++}
                                curentco[homes[home]] = crdn
                            }
                        } else {
                            time = times
                            Print("The fleet is NOT sended! "+err)
                            er = err
                            err = nil
                        }
                        if marker >= len(homes)-1 {err = er}
                    }
                    if slots == totalSlots {
                        home = len(homes)-1
                        time = times
                        fleetFlag = 2
                    }
                    Sleep(Random(1500, 3000))
                    if err != nil {slots = totalSlots}
                }
            } else {home = len(homes)-1}
            if PathfindersDebris == true {
                dflag = 0
                abr = 0
                nbr = 0
                curSystem = fromSystem
                if PathfinderSystemsRange == false {
                    curSystem = homeworld.GetCoordinate().System
                    toSystem = homeworld.GetCoordinate().System
                }
                for system = curSystem; system <= toSystem; system++ {
                    Sleep(Random(1000, 3000))
                    if GetSlots().InUse < totalUsl {
                        myShips, _ = homeworld.GetShips()
                        systemInfos, _ = GalaxyInfos(homeworld.GetCoordinate().Galaxy, system)
                        Dtarget, _ = ParseCoord(homeworld.GetCoordinate().Galaxy+":"+system+":"+16)
                        Debris, _ = ParseCoord("D:"+homeworld.GetCoordinate().Galaxy+":"+system+":"+16)
                        Print("Checking "+Dtarget)
                        if systemInfos.ExpeditionDebris.PathfindersNeeded >= Pnbr {
                            pp = systemInfos.ExpeditionDebris.PathfindersNeeded
                            if systemInfos.ExpeditionDebris.Metal == 0 && systemInfos.ExpeditionDebris.Crystal > 0 {Print("Found Crystal: "+systemInfos.ExpeditionDebris.Crystal)}
                            if systemInfos.ExpeditionDebris.Metal > 0 && systemInfos.ExpeditionDebris.Crystal == 0 {Print("Found Metal: "+systemInfos.ExpeditionDebris.Metal)}
                            if systemInfos.ExpeditionDebris.Metal > 0 && systemInfos.ExpeditionDebris.Crystal > 0 {Print("Found Metal: "+systemInfos.ExpeditionDebris.Metal+" and Crystal: "+systemInfos.ExpeditionDebris.Crystal)}
                            fleet, _ = GetFleets()
                            for f in fleet {
                                if f.Mission == RECYCLEDEBRISFIELD && f.ReturnFlight == false {
                                    if Debris == f.Destination {
                                        if f.Ships.Pathfinder < pp {
                                            abr = pp - f.Ships.Pathfinder
                                        } else {dflag = 1}
                                    }
                                }
                            }
                            if dflag == 0 {
                                f = NewFleet()
                                f.SetOrigin(homeworld)
                                f.SetDestination(Dtarget)
                                f.SetSpeed(HUNDRED_PERCENT)
                                f.SetMission(RECYCLEDEBRISFIELD)
                                if abr == 0 {
                                    nbr = systemInfos.ExpeditionDebris.PathfindersNeeded
                                } else {nbr = abr}
                                if nbr > myShips.Pathfinder {nbr = myShips.Pathfinder}
                                f.AddShips(PATHFINDER, nbr)
                                a, b = f.SendNow()
                                if b == nil {
                                    if nbr < systemInfos.ExpeditionDebris.PathfindersNeeded {Print("You don't have enough Ships for this debris field!")}
                                    if nbr > 1 {
                                        Print(nbr+" Pathfinders are sended successfully!")
                                    } else {Print(nbr+" Pathfinder is sended successfully!")}
                                } else {
                                    if nbr > 1 {
                                        Print("The Pathfinders are NOT sended! "+b)
                                    } else {Print("The Pathfinder is NOT sended! "+b)}
                                    break
                                }
                            } else {Print("Needed ships already are sended!")}
                        }
                    } else {
                        fleetFlag = 1
                        slots = totalSlots
                    }
                }
                if pp == 0 {Print("Not found any debris!")}
            }
            if cycle <= len(homes)-1 {cycle++}
            if home >= len(homes)-1 {
                for slots == totalSlots {
                    delay = Random(7*60, 12*60) // 7 - 12 minutes in seconds
                    if Repeat == true {
                        slots = GetSlots().ExpInUse
                        expslots = slots
                        if err != nil {
                            if slots != 0 {
                                for slots == expslots {
                                    delay = Random(7*60, 12*60) // 7 - 12 minutes in seconds
                                    Print("Please wait for the landing of your EXPO ships! Recheck after "+ShortDur(delay))
                                    Sleep(delay*1000)
                                    expslots = GetSlots().ExpInUse
                                    if slots > expslots {
                                        err = nil
                                        er = nil
                                    }
                                }
                            } else {
                                if cng == 0 {
                                    Print("All your EXPO ships are on the ground! Please, check your deuterium and make sure that you set the ships list correctly, then start the script again!")
                                    RepeatTimes = HowManyCycles
                                    useStartTime = false
                                    endFlag = 1
                                }
                            }
                        } else {
                            if fleetFlag == 1 {Print("All slots are busy now! Please, wait "+ShortDur(delay))}
                            if fleetFlag == 2 {Print("All EXPO slots are busy! Please, wait "+ShortDur(delay))}
                            Sleep(delay*1000)
                            slots = GetSlots().ExpInUse
                        }
                    } else {
                        slots = 1
                        totalSlots = 3
                    }
                }
                if RepeatTimes != HowManyCycles {
                    current = marker
                    if marker >= len(homes)-1 {
                        if HowManyCycles != false {
                            if Repeat == true {Print("You make full cycle of fleet sending "+RepeatTimes+"!")}
                            RepeatTimes++
                        }
                        current = -1
                        cng = 0
                        err = nil
                        er = nil
                    }
                    if Repeat == true {home = current}
                } else {
                    if endFlag == 0 {Print("You have reached the limit of repeats that you have set")}
                    Sleep(3000)
                }
            }
            Sleep(Random(1000, 3000))
        }
        if useStartTime == false {StopScript(__FILE__)}
    })
} else {
    Print("You typed wrong coordinates! - "+wrong)
    StopScript(__FILE__)
}
<-OnQuitCh
