package t

type Key string

const (
	NO         Key = "No"
	YES            = "Yes"
	CANCEL         = "Cancel"
	CANCELED       = "Canceled"
	COMPLETED      = "Completed"
	CONFIRM        = "Confirm"
	FAILURE        = "Failure"
	PROCESSING     = "Processing"
	WITHDRAW       = "Withdraw"
	ERROR          = "Error"
	CHECKING       = "Checking"
	TXPENDING      = "TxPending"
	TXCANCELED     = "TxCanceled"
	UNEXPECTED     = "Unexpected"

	CLAIMFAILED = "ClaimFailed"

	CALLBACKCOINFLIPWINNER = "CallbackCoinflipWinner"
	CALLBACKWINNER         = "CallbackWinner"
	CALLBACKERROR          = "CallbackError"
	CALLBACKEXPIRED        = "CallbackExpired"
	CALLBACKATTEMPT        = "CallbackAttempt"
	CALLBACKSENDING        = "CallbackSending"

	INLINEINVOICERESULT  = "InlineInvoiceResult"
	INLINEGIVEAWAYRESULT = "InlineGiveawayResult"
	INLINEGIVEFLIPRESULT = "InlineGiveflipResult"
	INLINECOINFLIPRESULT = "InlineCoinflipResult"
	INLINEHIDDENRESULT   = "InlineHiddenResult"

	LNURLINVALID = "LNURLInvalid"
	LNURLFAIL    = "LNURLFAIL"

	USERALLOWED       = "UserAllowed"
	SPAMFILTERMESSAGE = "SpamFilterMessage"

	PAYMENTFAILED       = "PaymentFailed"
	PAIDMESSAGE         = "PaidMessage"
	DBERROR             = "DBError"
	INSUFFICIENTBALANCE = "InsufficientBalance"
	TOOSMALLPAYMENT     = "TooSmallPayment"

	PAYMENTRECEIVED      = "PaymentReceived"
	FAILEDTOSAVERECEIVED = "FailedToSaveReceived"

	SPAMMYMSG    = "SpammyMsg"
	TICKETMSG    = "TicketMsg"
	FREEJOIN     = "FreeJoin"
	ASKTOCONFIRM = "AskToConfirm"

	HELPINTRO   = "HelpIntro"
	HELPSIMILAR = "HelpSimilar"
	HELPMETHOD  = "HelpMethod"

	RECEIVEHELP = "receiveHelp"

	PAYHELP = "payHelp"

	SENDHELP = "sendHelp"

	BALANCEHELP = "balanceHelp"

	GIVEAWAYHELP            = "giveawayHelp"
	GIVEAWAYMSG             = "GiveAwayMsg"
	GIVEAWAYCLAIM           = "GiveAwayClaim"
	GIVEAWAYSATSGIVENPUBLIC = "GiveawaySatsGivenPublic"

	COINFLIPHELP      = "coinflipHelp"
	COINFLIPWINNERMSG = "CoinflipWinnerMsg"
	COINFLIPGIVERMSG  = "CoinflipGiverMsg"
	COINFLIPAD        = "CoinflipAd"
	COINFLIPJOIN      = "CoinflipJoin"
	COINFLIPOVERQUOTA = "CoinflipOverQuota"
	COINFLIPRATELIMIT = "CoinflipRateLimit"

	GIVEFLIPHELP      = "giveflipHelp"
	GIVEFLIPMSG       = "GiveFlipMsg"
	GIVEFLIPWINNERMSG = "GiveflipWinnerMsg"
	GIVEFLIPAD        = "GiveflipAd"
	GIVEFLIPJOIN      = "GiveflipJoin"

	FUNDRAISEHELP        = "fundraiseHelp"
	FUNDRAISEAD          = "FundraiseAd"
	FUNDRAISEJOIN        = "FundraiseJoin"
	FUNDRAISECOMPLETE    = "FundraiseComplete"
	FUNDRAISERECEIVERMSG = "FundraiseReceiverMsg"
	FUNDRAISEGIVERMSG    = "FundraiseGiverMsg"

	BLUEWALLETHELP                = "bluewalletHelp"
	BLUEWALLETPASSWORDUPDATEERROR = "BluewalletPasswordUpdateError"
	BLUEWALLETCREDENTIALS         = "BluewalletCredentials"

	HIDEHELP             = "hideHelp"
	REVEALHELP           = "revealHelp"
	HIDDENREVEALBUTTON   = "HiddenRevealButton"
	HIDDENDEFAULTPREVIEW = "HiddenDefaultPreview"
	HIDDENWITHID         = "HiddenWithId"
	HIDDENSOURCEMSG      = "HiddenSourceMsg"
	HIDDENREVEALMSG      = "HiddenRevealMsg"
	HIDDENSTOREFAIL      = "HiddenStoreFail"
	HIDDENMSGNOTFOUND    = "HiddenMsgNotFound"
	HIDDENSHAREBTN       = "HiddenShareBtn"

	APPHELP = "appHelp"

	BITFLASHHELP         = "bitflashHelp"
	BITFLASHCONFIRM      = "BitflashConfirm"
	BITFLASHTXQUEUED     = "BitflashTxQueued"
	BITFLASHFAILEDTOSAVE = "BitflashFailedToSave"
	BITFLASHLIST         = "BitflashList"

	MICROBETHELP                = "microbetHelp"
	MICROBETBETHEADER           = "MicrobetBetHeader"
	MICROBETINVALIDRESPONSE     = "MicrobetInvalidResponse"
	MICROBETPAIDBUTNOTCONFIRMED = "MicrobetPaidButNotConfirmed"
	MICROBETPLACING             = "MicrobetPlacing"
	MICROBETPLACED              = "MicrobetPlaced"
	MICROBETFAILEDTOPAY         = "MicrobetFailedToPay"
	MICROBETLIST                = "MicrobetList"
	MICROBETBALANCEERROR        = "MicrobetBalanceError"
	MICROBETBALANCE             = "MicrobetBalance"

	SATELLITEHELP              = "satelliteHelp"
	SATELLITEFAILEDTOSTORE     = "SatelliteFailedToStore"
	SATELLITEFAILEDTOGET       = "SatelliteFailedToGet"
	SATELLITEPAID              = "SatellitePaid"
	SATELLITEFAILEDTOPAY       = "SatelliteFailedToPay"
	SATELLITEBUMPERROR         = "SatelliteBumpError"
	SATELLITEFAILEDTODELETE    = "SatelliteFailedToDelete"
	SATELLITEDELETEERROR       = "SatelliteDeleteError"
	SATELLITEDELETED           = "SatelliteDeleted"
	SATELLITETRANSMISSIONERROR = "SatelliteTransmissionError"
	SATELLITEQUEUEERROR        = "SatelliteQueueError"
	SATELLITEQUEUE             = "SatelliteQueue"
	SATELLITELIST              = "SatelliteList"

	GOLIGHTNINGHELP   = "golightningHelp"
	GOLIGHTNINGFAIL   = "GoLightningFail"
	GOLIGHTNINGFINISH = "GoLightningFinish"

	GIFTSHELP       = "giftsHelp"
	GIFTSERROR      = "GiftsError"
	GIFTSCREATED    = "GiftsCreated"
	GIFTSFAILEDSAVE = "GiftsFailedSave"
	GIFTSLIST       = "GiftsList"

	POKERHELP         = "pokerHelp"
	POKERDEPOSITFAIL  = "PokerDepositFail"
	POKERWITHDRAWFAIL = "PokerWithdrawFail"
	POKERBALANCEERROR = "PokerBalanceError"
	POKERSTATUS       = "PokerStatus"
	POKERNOTIFY       = "PokerNotify"
	POKERNOTIFYFRIEND = "PokerNotifyFriend"
	POKERSUBSCRIBED   = "PokerSubscribed"
	POKERSECRETURL    = "PokerSecretURL"
	POKERBALANCE      = "PokerBalance"

	PAYWALLHELP      = "paywallHelp"
	PAYWALLERROR     = "PaywallError"
	PAYWALLBALANCE   = "PaywallBalance"
	PAYWALLCREATED   = "PaywallCreated"
	PAYWALLLISTLINKS = "PaywallListLinks"

	ETLENEUMFAILEDTOPAY = "EtleneumFailedToPay"

	TOGGLEHELP = "toggleHelp"

	HELPHELP = "helpHelp"

	STOPHELP = "stopHelp"

	CONFIRMINVOICE     = "ConfirmInvoice"
	FAILEDDECODE       = "FailedDecode"
	NOINVOICE          = "NoInvoice"
	BALANCEMSG         = "BalanceMsg"
	FAILEDUSER         = "FailedUser"
	LOTTERYMSG         = "LotteryMsg"
	INVALIDPARTNUMBER  = "InvalidPartNumber"
	INVALIDAMOUNT      = "InvalidAmount"
	USERSENTTOUSER     = "UserSentToUser"
	USERSENTYOUSATS    = "UserSentYouSats"
	RECEIVEDSATSANON   = "ReceivedSatsAnon"
	FAILEDSEND         = "FailedSend"
	QRCODEFAIL         = "QRCodeFail"
	SAVERECEIVERFAIL   = "SaveReceiverFail"
	CANTSENDNORECEIVER = "CantSendNoReceiver"
	GIVERCANTJOIN      = "GiverCantJoin"
	CANTJOINTWICE      = "CantJoinTwice"
	CANTCANCEL         = "CantCancel"
	FAILEDINVOICE      = "FailedInvoice"
	ZEROAMOUNTINVOICE  = "ZeroAmountInvoice"
	INVALIDAMT         = "InvalidAmt"
	STOPNOTIFY         = "StopNotify"
	WELCOME            = "Welcome"
	WRONGCOMMAND       = "WrongCommand"
	RETRACTQUESTION    = "RetractQuestion"
	RECHECKPENDING     = "RecheckPending"
	TXNOTFOUND         = "TxNotFound"
	TXINFO             = "TxInfo"
	TXLIST             = "TxList"

	TUTORIALWALLET = "TutorialWallet"
	TUTORIALBLUE   = "TutorialBlue"
)
