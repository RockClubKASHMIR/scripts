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
 5. You can start the script at certain time. The end time is when repeats that you set for fleet send are completed.
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

