/***** This script is created by RockClubKASHMIR <discord @RockClubKASHMIR#8058> *****\
  discord channel for your personal orders and support - https://discord.gg/kbsdRCB
  
     DESCRIPTION
  1. Starts and stops the desired function or code you add at a certain time 
     Replace StartingAtsBot() and StoppingAtsBot() with your desired function or code, then click Save Script
*/

StartingAt = "10:23:39"
StoppingAt = "20:00:00"

CronExec(StartingAt, func() {
    StartingAtsBot()
    LogInfo("The function/or code starts at "+NowTimeString())
})

CronExec(StoppingAt, func() {
    StoppingAtsBot()
    LogInfo("The function/or code stops at "+NowTimeString())
})

<-OnQuitCh
