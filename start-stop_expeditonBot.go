/***** This script is created by RockClubKASHMIR <discord @RockClubKASHMIR#8058> *****\
  discord channel for your personal orders and support - https://discord.gg/kbsdRCB
*/

StartExpedition = "10:23:39"
StopExpedition = "20:00:00"

CronExec(StartExpedition, func() {
    StartExpeditionsBot()
    LogInfo("The expeditions Bot starts at "+NowTimeString())
})

CronExec(StopExpedition, func() {
    StopExpeditionsBot()
    LogInfo("The expeditions Bot stops at "+NowTimeString())
})

<-OnQuitCh
