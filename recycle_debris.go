/***** This script is created by RockClubKASHMIR <discord @RockClubKASHMIR#8058> *****\
discord channel for your personal orders and support - https://discord.gg/kbsdRCB

   v1.0
 
    DESCRIPTION
 1. This script always respect the reserved slots
 2. Checks for debris and recycle them
 3. You can set to recycle debris from more than 1 planet/moon
 4. Possibility to scan and recycle Debris in various solar systems
 5. You can set the minimum number of Recyclers needed for the debris you want to ignore
 6. Sends Recyclers to the same debris fields, only when ships are not enough to get them all
 7. Ability to set how many times to scan for debris
 8. Repeats the full scanning after the minutes that you set
 */

homes = ["M:1:2:3"] /* Replace M:1:2:3 with your coordinate - M for the moons, P for the planets.
  If you want to use more planets/moons, the list must look like this: homes = ["M:1:2:3", "M:2:2:3"] add as many planets/moons you want*/

SystemsRange = true // Do you want to check for debris in different solar systems? true = YES / false = NO
RangeRadius = 290  // Radius around your solar system. Set this if SystemsRange = true
Rnbr = 5  // The script will ignore debris less than for RECYCLERS that you set - The Maximum RECYCLERS is limited only of your RECYCLERS on the current moon/planet! You can set this value from 1 to the number you want

Repeat = true // Do you want to repeat the full scanning for debris? true = YES / false = NO
HowManyRepeats = 5 // Set how many times to repeat full scanning for debris - 0 means to repeat forever
PauseBetweenRepeats = 30 // Set the pause between repeats (in minutes)

//----- Please, don't change the code below -----\\
current = 0
wrong = []
homeworld = nil
PauseFarmingBot()
i = 0
ei = 0
er = nil
err = nil
nbr = 0
endFlag = 0
fleetFlag = 0
RepeatTimes = 1
if Rnbr < 1 {Rnbr = 1}
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
if HowManyRepeats == 0 {HowManyRepeats = false}
if homeworld != nil {
    ls = GetSlots()
    Sleep(2000)
    totalUsl = ls.Total - GetFleetSlotsReserved()
    for home = current; home <= len(homes)-1; home++ {
        i = 1
        pp = 0
        Dtarget = 0
        marker = home
        delay = 0
        ls = GetSlots()
        homeworld = GetCachedCelestial(homes[home])
        if homeworld.Coordinate.IsMoon() {
            Print("Your Moon is: "+homeworld.Coordinate)
        } else {Print("Your Planet is: "+homeworld.Coordinate)}
        fromSystem = homeworld.GetCoordinate().System - RangeRadius
        toSystem = homeworld.GetCoordinate().System + RangeRadius
        if fromSystem < 1 {fromSystem = 1}
        if toSystem > 499 {toSystem = 499}
        totalSlots = totalUsl
        slots = ls.InUse
        if slots < totalSlots {
            dflag = 0
            abr = 0
            curSystem = fromSystem
            if SystemsRange == false {
                curSystem = homeworld.GetCoordinate().System
                toSystem = homeworld.GetCoordinate().System
            }
            for system = curSystem; system <= toSystem; system++ {
                Sleep(Random(250, 1000))
                systemInfos, _ = GalaxyInfos(homeworld.GetCoordinate().Galaxy, system)
                planetInfo = systemInfos.Position(i)
                Debris, _ = ParseCoord("D:"+homeworld.GetCoordinate().Galaxy+":"+system+":"+i)
                if planetInfo != nil {
                    Print("Checking "+planetInfo.Coordinate)
                    if planetInfo.Debris.RecyclersNeeded >= Rnbr {
                        pp = planetInfo.Debris.RecyclersNeeded
                        if planetInfo.Debris.Metal == 0 && planetInfo.Debris.Crystal > 0 {Print("Found Crystal: "+planetInfo.Debris.Crystal)}
                        if planetInfo.Debris.Metal > 0 && planetInfo.Debris.Crystal == 0 {Print("Found Metal: "+planetInfo.Debris.Metal)}
                        if planetInfo.Debris.Metal > 0 && planetInfo.Debris.Crystal > 0 {Print("Found Metal: "+planetInfo.Debris.Metal+" and Crystal: "+planetInfo.Debris.Crystal)}
                        fleet, _ = GetFleets()
                        for f in fleet {
                            if f.Mission == RECYCLEDEBRISFIELD && f.ReturnFlight == false {
                                if Debris == f.Destination {
                                    if f.Ships.Recycler < pp {
                                        abr = pp - f.Ships.Recycler
                                    } else {
                                        dflag = 1
                                        nbr = 1
                                    }
                                }
                            }
                        }
                        ls = GetSlots()
                        Sleep(Random(250, 1400))
                        slots = ls.InUse
                        if slots < totalSlots {
                            if dflag == 0 {
                                myShips, _ = homeworld.GetShips()
                                f = NewFleet()
                                f.SetOrigin(homeworld)
                                f.SetDestination(planetInfo.Coordinate)
                                f.SetSpeed(HUNDRED_PERCENT)
                                f.SetMission(RECYCLEDEBRISFIELD)
                                if abr == 0 {
                                    nbr = planetInfo.Debris.RecyclersNeeded
                                } else {nbr = abr}
                                if nbr > myShips.Recycler {nbr = myShips.Recycler}
                                f.AddShips(RECYCLER, nbr)
                                a, err = f.SendNow()
                                if err == nil {
                                    slots = slots + 1
                                    if nbr < planetInfo.Debris.RecyclersNeeded {Print("You don't have enough Ships for this debris field!")}
                                    if nbr > 1 {
                                        Print(nbr+" Recyclers are sended successfully!")
                                    } else {Print(nbr+" Recycler is sended successfully!")}
                                } else {
                                    if nbr > 1 {
                                        Print("The Recyclers are NOT sended! "+err)
                                    } else {Print("The Recycler is NOT sended! "+err)}
                                    system = toSystem
                                    er = err
                                    err = nil
                                }
                            } else {Print("Needed ships already are sended!")}
                        }
                        if slots == totalSlots {
                            fleetFlag = 1
                            if system < toSystem {curSystem = system-1}
                            system = toSystem
                            current = marker-1
                            home = len(homes)-1
                        }
                    }
                }
                if i < 15 {
                    i++
                } else {i = 1}
                if marker >= len(homes)-1 {err = er}
            }
            if pp == 0 {Print("Not found any debris!")}
        } else {fleetFlag = 1}
        if err != nil {slots = totalSlots}
        if home >= len(homes)-1 {
            for slots == totalSlots {
                delay = Random(7*60, 13*60) // 7 - 13 minutes in seconds
                if Repeat == true {
                    if err != nil {
                        slots = GetSlots().InUse
                        expslots = slots
                        if slots > 0 {
                            Print("Please, wait until Recyclers returns! Re-check after "+ShortDur(delay))
                            Sleep(delay*1000)
                            expslots = GetSlots().InUse
                            if slots == expslots {slots = totalSlots}
                        }
                    } else {
                        if fleetFlag == 1 {Print("All slots are busy now! Please, wait "+ShortDur(delay))}
                        Sleep(delay*1000)
                        slots = GetSlots().InUse
                    }
                } else {
                    slots = 1
                    totalSlots = 3
                }
            }
            if RepeatTimes != HowManyRepeats {
                if err != nil || fleetFlag == 1 {
                    delay = 3
                    fleetFlag = 0
                } else {delay = Random((PauseBetweenRepeats-5)*60, PauseBetweenRepeats*60)}
                if marker >= len(homes)-1 {
                    if nbr == 0 {
                        Print("Not found any EXPEDITION debris!")
                        Sleep(3000)
                    }
                    if Repeat == true {
                        if RepeatTimes == 1 {Print("You have full scan for debris all coordinates "+RepeatTimes+" time")}
                        if RepeatTimes > 1 {Print("You have full scan for debris all coordinates "+RepeatTimes+" times")}
                    }
                    RepeatTimes++
                    Sleep(2000)
                    Print("Start searching for debris again after "+ShortDur(delay))
                    Sleep(delay*1000)
                    current = -1
                    err = nil
                    er = nil
                    nbr = 0
                }
                if Repeat == true {home = current}
            } else {
                Print("You have reached the limit of repeats that you have set")
                Sleep(3000)
            }
        }
        Sleep(Random(1000, 3000))
    }
} else {Print("You typed wrong coordinates! - "+wrong)}
