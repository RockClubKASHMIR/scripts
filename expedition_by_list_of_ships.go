/***** This script is created by RockClubKASHMIR <discord @RockClubKASHMIR#8058> *****\
 discord channel for your personal orders and support - https://discord.gg/kbsdRCB
 
    v6.0
 
    DESCRIPTION
 1. This script always respects the reserved slots
 2. Possibility to send fleets to EXPEDITION mission from more than 1 planet/moon
 3. You can create the list of ships by 2 methods (or by a combination of both of them):
    a. Automatic: All kind of ships with quantity 0 that you set will be sent dependent on free EXPEDITION slots (the full of ships with entered quantity 0 divided by the free EXPO slots)
       - if sentence = true, all kinds of ships set with quantity 0 will be sent as one fleet for each planet/moon.
    b. Manual: Enter the quantity of all kind of type of ships by yourself:
       - the ships set with this method will be accepted literally, and if any kind type of your ships is even 1 less than the quantity you set, the fleet will not be sent!
 4. Possibility to send the fleets on EXPEDITION mission to the radius of solar systems entered by you around your current solar system or only to your current solar system
 5. Evenly distribution of EXPEDITION slots per each moon/planet or use all EXPEDITION slots per every planet/moon
 6. Check for EXPEDITION Debris and recycle them (if you are Discoverer and have Pathfinders)
 7. Possibility to make a scan and recycle debris at a range of your solar system or to your solar system only
 8. You can set a minimum amount of pathfinders for recycling
 9. Sends Pathfinders to same debris more than once only if already sent ships are not enough to get all resources
10. Possibility to repeat the sending of EXPEDITION fleets many times - you can set how many
11. You can start this script at a specific time. Sending of the fleets will stop after the number of repeats that you set
*/

homes = ["M:2:205:8", "M:4:182:5"] // Replace coordinates M:1:2:3 with your coordinates. "M" means moon, "P" means planet.
// You can add as many planets/moons you want. The list of planets/moons should look like this: homes = ["M:1:2:3", "M:2:2:3"]

shipsList = {LARGECARGO: 0, LIGHTFIGHTER: 0, PATHFINDER: 100}// Set your list of the ships

RangeRadius = 5  // Enter the radius of systems around your current system you want to send your fleets (only if you planning to use this option) 
SystemsRange = false // Do you want to send the fleets sent on the EXPEDITION mission in a radius of systems around your current solar system? true = YES / false = NO
sendWhenFleetBack = false // Do you want to wait for the return of all fleets sent on the EXPEDITION mission before sending them all again? true = YES / false = NO
sendAtOnce = false // Do you want to send all kinds of ships with quantity 0 as one fleet for each your planet/moon? true = YES / false = NO

DurationOfExpedition = 1 // Enter duration (in hours) of the EXPEDITION: minimum 1 - maximum 8 hours
Repeat = true // Do you want to repeat the full cycle of fleet sending? true = YES / false = NO
HowManyCycles = 5 // Set the limit of repeats of whole cycle of EXPO fleet sending - 0 means forewer

PathfindersDebris = true // Do you want to detect/get debris from EXPEDITION? true = YES / false = NO
Pnbr = 5  // Enter the minimum quantity of Pathfinders that can take the desired minimum amount of EXPEDITION debris

myTime = "09:33:00"// Enter the time of starting the fleets sending; Hour: 00 - 23, Minute: 00 - 59, Seconds: 00 - 59
useStartTime = false // Do you want to run this script at a specific time every day? true = YES / false = NO

//----- Please, don't change the code below -----\\
current = 0
wrong = []
curentco = {}
waves = {}
homeworld = nil
PauseFarmingBot()
StopHunter()
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
calc = 0
if (Pnbr < 1) {Pnbr = 1}
for home in homes {
    flag = 1
    hh, _ = ParseCoord(home)
    for celestial in GetCachedCelestials() {
        if celestial.Coordinate == hh {
            ei++
            flag = 0
        }
    }
    if flag == 1 {wrong += home}
    i++
}
if ei == len(homes) {homeworld = GetCachedCelestial(homes[0])}
if len(shipsList) > 0 {
    for ShipID, num in shipsList {
        if num == 0 {calc = 1}
    }
} else {
    Print("Your Ship's list is emty!")
    StopScript(__FILE__)
}
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
        ls = GetSlots()
        Sleep(2000)
        totalUsl = ls.Total - GetFleetSlotsReserved()
        totalExpSlots = ls.ExpTotal
        for home = current; home <= len(homes)-1; home++ {
            trgList = []
            ttargets = {}
            pp = 0
            cs = 0
            Dtarget = 0
            marker = home
            rcyc = 0
            ls = GetSlots()
            homeworld = GetCachedCelestial(homes[home])
            if homeworld.Coordinate.IsMoon() {
                Print("Your Moon is: "+homeworld.Coordinate)
            } else {Print("Your Planet is: "+homeworld.Coordinate)}
            trgList = GetSystemsInRange(homeworld.GetCoordinate().System, RangeRadius)
            crdn = trgList[0]
            ExpsTemp = 0
            if SystemsRange == true && cycle >= len(homes)-1 {
                for id, num in curentco {
                    if id == homes[home] {
                        crdn = num
                        for n in trgList {
                            if crdn != n {cs++}
                        }
                    }
                }
            }
            totalSlots = totalUsl
            slots = ls.InUse
            bk = 0
            currentTime = bk
            times = totalExpSlots
            if slots < totalSlots {
                slots = ls.ExpInUse
                totalSlots = totalExpSlots
                ExpsTemp = 1
                if slots == totalSlots {fleetFlag = 2}
            } else {fleetFlag = 1}
            if err != nil {slots = totalSlots}
            if slots < totalSlots {
                Expos = totalExpSlots - slots
                if sendWhenFleetBack == false {
                    slotMarker = totalExpSlots-marker
                    times = slotMarker/len(homes)
                    if times > Floor(times) {times = Floor(times) + 1}
                    if times < 1 {times = 1}
                }
                Flts, _ = GetFleets()
                bk = 0
                for f in Flts {
                    if f.Mission == EXPEDITION {
                        hh, _ = ParseCoord(homes[home])
                        if hh == f.Origin {bk = bk + 1}
                    }
                }
                currentTime = bk
                if sendAtOnce == true {
                    times = 1
                    Expos = 1
                }
                Expos = times - bk
                if Expos <= 0 {
                    currentTime = times
                    Print("There are no EXPO fleets to send here!")
                } else {Print(Expos+" slots will be used")}
                
                for time = currentTime; time < times; time++ {
                    myShips, _ = homeworld.GetShips()
                    tt = 0
                    rtt = 0
                    ExpFleet = {}
                    if ExpsTemp == 0 {
                        totalSlots = totalUsl
                        ls = GetSlots()
                        slots = ls.InUse
                        Sleep(800)
                        if slots < totalSlots {
                            slots = ls.ExpInUse
                            Sleep(800)
                            totalSlots = totalExpSlots
                            if slots == totalSlots {fleetFlag = 2}
                        } else {fleetFlag = 1}
                    }
                    if err != nil {slots = totalSlots}
                    if slots < totalSlots {
                        ExpsTemp == 0
                        if SystemsRange == false {
                            Dtarget, _ = ParseCoord(homeworld.GetCoordinate().Galaxy+":"+homeworld.GetCoordinate().System+":"+16)
                        }
                        if SystemsRange == true {
                            if cs > len(trgList)-1 {
                                crdn = trgList[0]
                                cs = 0
                            }
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
                                            temp = (num/100)*40
                                            if myShips.ByID(ShipID) < num && myShips.ByID(ShipID) >= temp {num = myShips.ByID(ShipID)}
                                            if myShips.ByID(ShipID) < num && myShips.ByID(ShipID) < temp {num = 0}
                                        }
                                        if sendAtOnce == true {num = myShips.ByID(ShipID)}
                                        if num < 1 {num = 0}
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
                                        }
                                        if ShipID == PATHFINDER {
                                            if myShips.ByID(ShipID) >= num {
                                                ExpFleet[ShipID] = num
                                                tt = tt + 1
                                            }
                                            if len(shipsList) > 1 && myShips.ByID(ShipID) < num {
                                                num = myShips.ByID(ShipID)
                                                ExpFleet[ShipID] = num
                                                tt = tt + 1
                                            }
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
                            waves[homes[home]] = 1
                            Print(explist+" are successfully sent to "+Dtarget)
                            if SystemsRange == true {
                                if cs <= len(trgList)-1 {
                                    cs++
                                    if cs > len(trgList)-1 {cs = 0}
                                    crdn = trgList[cs]
                                }
                                curentco[homes[home]] = crdn
                            }
                            ttargets[Dtarget] = homeworld.Coordinate
                        } else {
                            time = times
                            Print("The fleet is NOT sended! "+err)
                            er = err
                            err = nil
                        }
                        if marker >= len(homes)-1 {err = er}
                        
                    }
                    if slots == totalSlots && err == nil {
                        time = times
                        fleetFlag = 2
                    }
                    if err != nil {slots = totalSlots}
                    Sleep(Random(1500, 3000))
                }
            } else {home = len(homes)-1}
            func sendPathfinders() {
                if PathfindersDebris == true {
                    pp = 0
                    for Dtarget, homewr in ttargets {
                        homeworld = GetCachedCelestial(homewr)
                        dflag = 0
                        abr = 0
                        nbr = 0
                        Sleep(Random(300, 850))
                        systemInfos, _ = GalaxyInfos(Dtarget.Galaxy, Dtarget.System)
                        Debris, _ = ParseCoord("D:"+Dtarget.Galaxy+":"+Dtarget.System+":"+16)
                        //Print("Checking "+Dtarget)
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
                            ls = GetSlots()
                            Sleep(Random(1000, 3000))
                            aaz = ls.InUse
                            if aaz < totalUsl {
                                if dflag == 0 {
                                    myShips, _ = homeworld.GetShips()
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
                                        aaz = aaz + 1
                                        if aaz == totalUsl {
                                            fleetFlag = 1
                                            system = toSystem
                                        }
                                        if nbr < systemInfos.ExpeditionDebris.PathfindersNeeded {Print("You don't have enough Ships for this debris field!")}
                                        if nbr > 1 {
                                            Print(nbr+" Pathfinders are successfully sent!")
                                        } else {Print(nbr+" Pathfinder is successfully sent!")}
                                    } else {
                                        if nbr > 1 {
                                            Print("The Pathfinders have NOT been sent! "+b)
                                        } else {Print("The Pathfinder has NOT been sent!"+b)}
                                        break
                                    }
                                }
                            } else {fleetFlag = 1}
                        }
                    }
                    if pp == 0 {Print("Not found any debris!")}
                }
            }
            if cycle <= len(homes)-1 {cycle++}
            ls = GetSlots()
            if ls.InUse == totalUsl || ls.ExpInUse == totalExpSlots {
                if sendWhenFleetBack == true && sendAtOnce == false {
                    if ls.ExpInUse < totalExpSlots {fleetFlag = 0}
                }
                home = len(homes)-1
                slots = totalSlots
                current = marker
            }
            if sendAtOnce == true {
                if marker >= len(homes)-1 {
                    if ls.ExpInUse < totalExpSlots && err == nil {
                        err = "no ships to send"
                        slots = totalSlots
                    }
                }
            }
            if home >= len(homes)-1 {
                for slots == totalSlots {
                    delay = 8000000
                    totaltime = delay
                    ttargets = {}
                    fleet, _ = GetFleets()
                    for f in fleet {
                        u = 0
                        if f.Mission == EXPEDITION {
                            if f.ReturnFlight == true {
                                ttargets[f.Destination] = f.Origin
                                totaltime = f.BackIn
                                delay = 1
                            }
                            if f.ReturnFlight == false {u = (f.ArriveIn+4)+3600*DurationOfExpedition}
                            if u < 0 {u = u*-1}
                            if u > 0 {
                                if delay > u {delay = u}
                            }
                        }
                    }
                    if rcyc == 1 {
                        Print("Preparing to check for debris...")
                        Sleep(3000)
                        sendPathfinders()
                        if fleetFlag == 1 {
                            err = nil
                            er = nil
                        }
                    }
                    if Repeat == true {
                        if err != nil {
                            slots = GetSlots().ExpInUse
                            expslots = slots
                            if slots > 0 {
                                Print("Please wait for the landing of the fleets sent on EXPEDITION! Please wait "+ShortDur(delay))
                                Sleep(delay*1000)
                                expslots = GetSlots().ExpInUse
                                if slots > expslots {
                                    err = nil
                                    er = nil
                                } else {slots = totalSlots}
                            } else {
                                if cng == 0 {
                                    Print("All your ships from the list are on the ground! Please, check your deuterium and make sure that the list of ships is correctly set, then start the script again!")
                                    RepeatTimes = HowManyCycles
                                    useStartTime = false
                                    endFlag = 1
                                }
                            }
                        } else {
                            if rcyc == 1 {delay = totaltime}
                            if fleetFlag == 0 {Print("Please, wait till all your fleets sent on EXPEDITION arrives! Please, wait "+ShortDur(delay))}
                            if fleetFlag == 1 {Print("The maximum number of flights has been reached! Please, wait "+ShortDur(delay))}
                            if fleetFlag == 2 {
                                if delay == 1 {Print("All slots for Expedition are busy!")}
                                if delay > 1 {Print("All slots for Expedition are busy! Please, wait "+ShortDur(delay))}
                            }
                            Sleep(delay*1000)
                            if fleetFlag == 1 {slots = GetSlots().InUse}
                            if  fleetFlag == 0 || fleetFlag == 2 {
                                slots = GetSlots().ExpInUse
                                if sendWhenFleetBack == true && slots >= 1 {
                                    if slots < totalSlots {
                                        fleetFlag = 0
                                        slots = totalSlots
                                    }
                                }
                            }
                        }
                    } else {
                        slots = 1
                        totalSlots = 3
                    }
                    rcyc = 1
                }
                if RepeatTimes != HowManyCycles {
                    if marker >= len(homes)-1 {
                        if len(waves) == len(homes) {
                            if HowManyCycles != false {
                                if Repeat == true {Print("You make full cycle of fleet sending "+RepeatTimes+"!")}
                                RepeatTimes++
                                waves = {}
                            }
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
