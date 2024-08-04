package address

// ###########################################################//

var (
	SYSCON_SYSCLK_CONF_REG    = AddressCellObj{0x0000, 0x0004}     //Configures system clock frequency
	SYSCON_XTAL_TICK_CONF_REG = AddressCellObj{0x0004, 0x0008}     //Configures the divider value of REF_TICK
	SYSCON_PLL_TICK_CONF_REG  = AddressCellObj{0x0008, 0x000C}     //Configures the divider value of REF_TICK
	SYSCON_CK8M_TICK_CONF_REG = AddressCellObj{0x000C, 0x003C}     //Configures the divider value of REF_TICK
	SYSCON_APLL_TICK_CONF_REG = AddressCellObj{0x003C, 0x007C}     //Configures the divider value of REF_TICK
	SYSCON_DATE_REG           = AddressCellObj{0x007C, 0x007C + 5} //Chip revision register
)
