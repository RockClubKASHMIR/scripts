//==== This script is created by RockClubKASHMIR ====

/* DESCRIPTION
   This script find automatically your planet/moon with highgly amount of Pathfinders!
  
  ONLY if my automatic method of finding your moon/planet not satisfied you;
  - replace all rows between // START and // END with origin = GetCachedCelestial("M:1:2:3") where on "M:1:2:3" must type your coordinate - M for the moon, P for planet
*/

fromSystem = 1 // Your can change this value as you want
toSystem = 200 // Your can change this value as you want
Pnbr = 2  // When Pnbr = 2, the script will search debris for minimum 2 Pathfinders, but this is NOT mean that this is a Limit for Maximum Pathfinders! You can set this value from 0, to the number you want

//----
curSystem = fromSystem
origin = nil
flts = 0
nbr = 0
err = nil
if (Pnbr < 0) {Pnbr = 0}
totalSlots = GetSlots().Total - GetFleetSlotsReserved()
// START
for celestial in GetCachedCelestials() {
    ships, _ = celestial.GetShips()
    if ships.Pathfinder > flts {
        flts = ships.Pathfinder
        origin = celestial 
    }
}
// END
if origin != nil {
    if IsDiscoverer() {
        Print("Your origin is "+origin.Coordinate)
        if toSystem > 499 || toSystem == 0 {toSystem = -1}
        if fromSystem > toSystem {Print("Please, type correctly fromSystem and/or toSystem!")}
        for system = curSystem; system <= toSystem; system++ {
            pp = 0
            dflag = 0
            abr = 0
            nbr = 0
            systemInfos, b = GalaxyInfos(origin.GetCoordinate().Galaxy, system)
            Dtarget, _ = ParseCoord(origin.GetCoordinate().Galaxy+":"+system+":"+16)
            Debris, _ = ParseCoord("D:"+origin.GetCoordinate().Galaxy+":"+system+":"+16)
            Sleep(Random(500, 1500)) // for avoid ban
            slots = GetSlots().InUse
            if err != nil {slots = totalSlots}
            if slots < totalSlots {
                if b == nil {
                    Print("Checking "+Dtarget)
                    if systemInfos.ExpeditionDebris.PathfindersNeeded >= Pnbr { 
                        ships, _ = origin.GetShips()
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
                            f.SetOrigin(origin)
                            f.SetDestination(Dtarget)
                            f.SetSpeed(HUNDRED_PERCENT)
                            f.SetMission(RECYCLEDEBRISFIELD)
                            if abr == 0 {
                                nbr = systemInfos.ExpeditionDebris.PathfindersNeeded
                            } else {nbr = abr}
                            if nbr > ships.Pathfinder {nbr = ships.Pathfinder}
                            f.AddShips(PATHFINDER, nbr)
                            a, err = f.SendNow()
                            if err == nil {
                                if nbr < systemInfos.ExpeditionDebris.PathfindersNeeded {Print("You don't have enough Ships for this debris field!")}
                                if nbr > 1 {
                                    Print(nbr+" Pathfinders are sended successfully!")
                                } else {Print(nbr+" Pathfinder is sended successfully!")}
                            } else {
                                if nbr > 1 {
                                    Print("The Pathfinders are NOT sended! "+err)
                                } else {Print("The Pathfinder is NOT sended! "+err)}
                            }
                        } else {Print("Needed ships already are sended!")}
                    }
                }
            } else {
                for slots == totalSlots {
                    if err != 0 {
                        Print("Please wait till ships lands! Recheck after "+ShortDur(4*60))
                        Sleep(4*60*1000)
                        ships, _ = origin.GetShips()
                        if ships.Pathfinder > 0 {slots = GetSlots().InUse}
                        err = nil
                    } else {
                        Print("All Fleet slots are busy now! Please, wait "+ShortDur(4*60))
                        Sleep(4*60*1000)
                        slots = GetSlots().InUse
                    }
                    curSystem = system-1
                }
            }
            if b == nil {
                if system >= toSystem {
                    delay = Random(50*60, 90*60)
                    sleepDelay = delay*1000
                    Print("Will Start searching again after "+ShortDur(delay))
                    Sleep(sleepDelay)
                    Print("Start searching again...")
                    curSystem = fromSystem-1
                    system = curSystem
                }
            } else {
                Print("Please, type correctly fromSystem and/or toSystem!")
                break
            }
        }
    } else {Print("You are not DISCOVERER!")}
} else {Print("You don't have Pathfinders on your Planets/Moons!")}
