//==== This script is created by RockClubKASHMIR ====

//--- WARNING!!! This script can work ONLY if you are Discoverer! ---
fromSystem = 1 // Your can change this value as you want
toSystem = 200 // Your can change this value as you want
Pnbr = 1  // When Pnbr = 1, the script will search debris for minimum 2 Pathfinders. You can set this value from 0, to the number you want
times = 1 // if times = 1, the script will full scan 2 times the entire galaxy, from system, to system you set. You can set this value from 0, to the number you want

//----
cycle = 0
curSystem = fromSystem
origin = nil
flts = 0
nbr = 0
err = nil
if (Pnbr < 0) {Pnbr = 0}
if (times < 0) {times = 0}
totalSlots = GetSlots().Total - GetFleetSlotsReserved()
// Start to search highest amount of Pathfinders on all your Planets and Moons(if you have some)
for celestial in GetCachedCelestials() {
    ships, _ = celestial.GetShips()
    if ships.Pathfinder > flts {
        flts = ships.Pathfinder
        origin = celestial // Your Planet(or Moon), with more Pathfinders
    }
}
if origin != nil {
    Print("Your origin is "+origin.Coordinate)
    if toSystem > 499 || toSystem == 0 {toSystem = -1}
    if fromSystem > toSystem {Print("Please, type correctly fromSystem and/or toSystem!")}
    for system = curSystem; system <= toSystem; system++ {
        systemInfos, b = GalaxyInfos(origin.GetCoordinate().Galaxy, system)
        Dtarget, _ = ParseCoord(origin.GetCoordinate().Galaxy+":"+system+":"+16)
        Sleep(Random(500, 1500)) // for avoid ban
        slots = GetSlots().InUse
        if err != nil {slots = totalSlots}
        if slots < totalSlots {
            if b == nil {
                Print("Checking "+Dtarget)
                if systemInfos.ExpeditionDebris.PathfindersNeeded > Pnbr { 
                    ships, _ = origin.GetShips()
                    if systemInfos.ExpeditionDebris.Metal == 0 && systemInfos.ExpeditionDebris.Crystal > 0 {Print("Found Crystal: "+systemInfos.ExpeditionDebris.Crystal)}
                    if systemInfos.ExpeditionDebris.Metal > 0 && systemInfos.ExpeditionDebris.Crystal == 0 {Print("Found Metal: "+systemInfos.ExpeditionDebris.Metal)}
                    if systemInfos.ExpeditionDebris.Metal > 0 && systemInfos.ExpeditionDebris.Crystal > 0 {Print("Found Metal: "+systemInfos.ExpeditionDebris.Metal+" and Crystal: "+systemInfos.ExpeditionDebris.Crystal)}
                    f = NewFleet()
                    f.SetOrigin(origin)
                    f.SetDestination(Dtarget)
                    f.SetSpeed(HUNDRED_PERCENT)
                    f.SetMission(RECYCLEDEBRISFIELD)
                    nbr = systemInfos.ExpeditionDebris.PathfindersNeeded
                    if systemInfos.ExpeditionDebris.PathfindersNeeded > ships.Pathfinder {nbr = ships.Pathfinder}
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
                            SendTelegram(TELEGRAM_CHAT_ID, "The Pathfinders are NOT sended! "+err)
                        } else {
                            Print("The Pathfinder is NOT sended! "+err)
                            SendTelegram(TELEGRAM_CHAT_ID, "The Pathfinder is NOT sended! "+err)
                        }
                    }
                }
            }
        } else {
            for slots == totalSlots {
                if err != 0 {
                    Print("Please wait till ships lands! Recheck after "+ShortDur(2*60))
                    Sleep(2*60*1000)
                    ships, _ = origin.GetShips()
                    if ships.Pathfinder > 0 {slots = GetSlots().InUse}
                    err = nil
                } else {
                    Print("All Fleet slots are busy now! Please, wait "+ShortDur(2*60))
                    Sleep(2*60*1000)
                    slots = GetSlots().InUse
                }
                curSystem = system-1
            }
        }
        if b == nil {
            if system >= toSystem {
                if times > 0 {
                    if cycle < times {
                        cycle++
                        if nbr == 0 {Print("Not found any debris! Start searching again...")}
                        curSystem = fromSystem-1
                        system = curSystem
                        Sleep(4000)
                    } else {
                        Print("You made "+(times+1)+" times full scan all systems chosen by you! The script turns off")
                        SendTelegram(TELEGRAM_CHAT_ID, "You made "+(times+1)+" times full scan all systems chosen by you! The script turns off")
                        break
                    }
                } else {
                    Print("You made full scan all systems chosen by you! The script turns off")
                    SendTelegram(TELEGRAM_CHAT_ID, "You made full scan all systems chosen by you! The script turns off")
                    break
                }
            }
        } else {
            Print("Please, type correctly fromSystem and/or toSystem!")
            break
        }
    }
} else {Print("You don't have Pathfinders on your Planets/Moons!")}
