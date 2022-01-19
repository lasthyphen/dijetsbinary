// (c) 2019-2020, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package genesis

import (
	"time"

	"github.com/ava-labs/avalanchego/utils/units"
)

var (
	mainnetGenesisConfigJSON = `{
		"networkID": 1,
		"allocations": [
			{
				"ethAddr": "0x828c2d1ec828b6b366656652ff6ee1af487c9b02",
				"djtxAddr": "X-dijets18umxudj8hxry68qz2gas40yy6rg4safw2r5wlv",
				"initialAmount": 500000000000000,
				"unlockSchedule": [
					{
						"amount": 4375000000000000,
						"locktime": 1637263338
					}
				]
			},
			{
				"ethAddr": "0xbcad856abde568fabdfafae989c902e9739a6ab7",
				"djtxAddr": "X-dijets1t7693hu760mdtce2e6jecmhd32rurh8ey0jl8j",
				"initialAmount":10000000000000000,
				"unlockSchedule": [
					{
						"amount": 4375000000000000,
						"locktime": 1637263338
					}
				]
			},
			{
				"ethAddr": "0xc5f2d024e2995acf577fa9f2c18d96ca2d915aec",
				"djtxAddr": "X-dijets1pr5rxgtpu98yuk3hvzawj526u8ed8fjtqnkttn",
				"initialAmount": 500000000000000,
				"unlockSchedule": [
					{
						"amount": 4375000000000000,
						"locktime": 1637868138
					}
				]
			},
			{
				"ethAddr": "0xc4be9c2ef30f64930b643bec3f36c88540670094",
				"djtxAddr": "X-dijets1l6u9wyrk7ald0qt7xvs6ck9j7628q6pu4edtqh",
				"initialAmount": 500000000000000,
				"unlockSchedule": [
					{
						"amount": 4375000000000000,
						"locktime": 1638300138
					}
				]
			},
			{
				"ethAddr": "0xf0a343fce12db62792bbadf8d596ee5b9077cee9",
				"djtxAddr": "X-dijets18cyk5dhgj6vp40urrrs7e9k5dfhkm5netm4rj0",
				"initialAmount": 500000000000000,
				"unlockSchedule": [
					{
						"amount": 4375000000000000,
						"locktime": 1638818538
					}
				]
			},
			{
				"ethAddr": "0x14a12e60daefa8969bda45003e520843213838a3",
				"djtxAddr": "X-dijets1x8ra9hrcy59sl7sfrn5dphww0hhv2kt73lrhfp",
				"initialAmount": 500000000000000,
				"unlockSchedule": [
					{
						"amount": 4375000000000000,
						"locktime": 1639423338
					}
				]
			},
			{
				"ethAddr": "0x30efe1f1fcf6630bcde520fba38b4db88bf63051",
				"djtxAddr": "X-dijets1z688szk5pdkj2sh7w3gnpmdep0gu96nyjqj99a",
				"initialAmount": 500000000000000,
				"unlockSchedule": [
					{
						"amount": 4375000000000000,
						"locktime": 1640028138
					}
				]
			},
			{
				"ethAddr": "0x3686ade70a6cddce02b2bd7c0e9c53d6c0b96329",
				"djtxAddr": "X-dijets1y4m6490pt4r7kaglegx2s8yesktcehgzgcw502",
				"initialAmount": 500000000000000,
				"unlockSchedule": [
					{
						"amount": 4375000000000000,
						"locktime": 1640632938
					}
				]
			},
			{
				"ethAddr": "0x581ba4548ba11777281d51a383fc0b9bb1eda4e6",
				"djtxAddr": "X-dijets1ser694s6c9tejs6xgu6cj8vp73ktskxmxrzd2r",
				"initialAmount": 500000000000000,
				"unlockSchedule": [
					{
						"amount": 4375000000000000,
						"locktime": 1640892138
					}
				]
			}
		],
		"startTime": 1636700400,
		"initialStakeDuration": 31536000,
		"initialStakeDurationOffset": 5400,
		"initialStakedFunds": [
			"X-dijets1pr5rxgtpu98yuk3hvzawj526u8ed8fjtqnkttn"
		],
		"initialStakers": [
			{
				"nodeID": "NodeID-7Xhw2mDxuDS44j42TCB6U5579esbSt3Lg",
				"rewardAddress": "X-dijets1pr5rxgtpu98yuk3hvzawj526u8ed8fjtqnkttn",
				"delegationFee": 1000000
			},
			{
				"nodeID": "NodeID-MFrZFVCXPv5iCn6M9K6XduxGTYp891xXZ",
				"rewardAddress": "X-dijets1pr5rxgtpu98yuk3hvzawj526u8ed8fjtqnkttn",
				"delegationFee": 500000
			},
			{
				"nodeID": "NodeID-NFBbbJ4qCmNaCzeW7sxErhvWqvEQMnYcN",
				"rewardAddress": "X-dijets1pr5rxgtpu98yuk3hvzawj526u8ed8fjtqnkttn",
				"delegationFee": 250000
			},
			{
				"nodeID": "NodeID-GWPcbFJZFfZreETSoWjPimr846mXEKCtu",
				"rewardAddress": "X-dijets1pr5rxgtpu98yuk3hvzawj526u8ed8fjtqnkttn",
				"delegationFee": 125000
			},
			{
				"nodeID": "NodeID-P7oB2McjBGgW2NXXWVYjV8JEDFoW9xDE5",
				"rewardAddress": "X-dijets1pr5rxgtpu98yuk3hvzawj526u8ed8fjtqnkttn",
				"delegationFee": 62500
			}
		],
		"cChainGenesis": "{\"config\":{\"chainId\":43114,\"homesteadBlock\":0,\"daoForkBlock\":0,\"daoForkSupport\":true,\"eip150Block\":0,\"eip150Hash\":\"0x2086799aeebeae135c246c65021c82b4e15a2c451340993aacfd2751886514f0\",\"eip155Block\":0,\"eip158Block\":0,\"byzantiumBlock\":0,\"constantinopleBlock\":0,\"petersburgBlock\":0,\"istanbulBlock\":0,\"muirGlacierBlock\":0,\"apricotPhase1BlockTimestamp\":0,\"apricotPhase2BlockTimestamp\":0},\"nonce\":\"0x0\",\"timestamp\":\"0x0\",\"extraData\":\"0x00\",\"gasLimit\":\"0x5f5e100\",\"difficulty\":\"0x0\",\"mixHash\":\"0x0000000000000000000000000000000000000000000000000000000000000000\",\"coinbase\":\"0x0000000000000000000000000000000000000000\",\"alloc\":{\"8db97C7cEcE249c2b98bDC0226Cc4C2A57BF52FC\":{\"balance\":\"0x295BE96E64066972000000\"}},\"number\":\"0x0\",\"gasUsed\":\"0x0\",\"parentHash\":\"0x0000000000000000000000000000000000000000000000000000000000000000\"}",
		"message": "{{ fun_quote }}"
	}`

	// MainnetParams are the params used for mainnet
	MainnetParams = Params{
		TxFeeConfig: TxFeeConfig{
			TxFee:                 units.MilliDjtx,
			CreateAssetTxFee:      units.MilliDjtx,
			CreateSubnetTxFee:     100 * units.MilliDjtx,
			CreateBlockchainTxFee: 100 * units.MilliDjtx,
		},
		StakingConfig: StakingConfig{
			UptimeRequirement:  .8, // 80%
			MinValidatorStake:  2 * units.KiloDjtx,
			MaxValidatorStake:  3 * units.MegaDjtx,
			MinDelegatorStake:  25 * units.Djtx,
			MinDelegationFee:   20000, // 2%
			MinStakeDuration:   24 * time.Hour,
			MaxStakeDuration:   365 * 24 * time.Hour,
			StakeMintingPeriod: 365 * 24 * time.Hour,
		},
	}
)
