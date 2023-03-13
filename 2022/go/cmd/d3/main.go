package main

import (
	"fmt"
	"strings"
)

func getInput() string {
	// return `vJrwpWtwJgWrhcsFMMfFFhFp
	// jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
	// PmmdzqPrVvPwwTWBwg
	// wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
	// ttgJtRGJQctTZtZT
	// CrZsJsPPZsGzwwsLwLmpwMDw`
	return `zBBtHnnHtwwHplmlRlzPLCpp
	vvhJccJFGFcNsdNNJbhJsJQplQMRLQMlfdfTPCLfQQCT
	GPhjcjhZDjWtnSVH
	BNhHVhrGNVTbDHdDJdJRPJdSQQSJwPjR
	lvtsfbsqzwSnJcvjSm
	MftttFLftZMLgtgMbltMqZzbDNrTpVGhNWrDTrpTGNpZGZhD
	VSSHcTgTtTdtllZlzmmbljTn
	RqMqsFfQLLFLQFMMfRLPZLvPpCfWrbpmCbjCnfjlWmnrmmnm
	hqRDqPDRsqNHwtHSNBZtJd
	tNFDpDFrtdjfmjjjFmFFdScpZhZScTJgpHccHhMJgS
	lLzSlSCQqbsVhBghggBZgCcJ
	zRLVVLQnvQqVVzRldfWrwffjjdwSdfjv
	bpWqqqWvHBpwGBCCRl
	hJdjdJFQqdBBDMMC
	tFFzJZFtJSqtZJQsWLbNSTnffHfvTH
	lFhRZhFjPlqMlJqZJlJcRLwrLrwStRwtsVVtVSrgRV
	WcpDvDfBmpDHzWBDbpbmWmNVSSTzLTtrVswgttVVzwwr
	pbWfmGBpHfDmWnvvGbmWnjjMqPJMlMFPdGcjqPqPhP
	NjFNRlpVLFCSSlbBWWfw
	pssPZQQsMnzmtnQPttzDBbBJBcrrJWbrZSBJSbfC
	QTHPHspMNGHdhvRR
	QfPdSJfFJmthSthtwbsNLbPLlLTLpbvP
	nHnMBnZqqgBMnWrZMqnZVcbCqRwNsvblRwppbllTsRNp
	nZHBHznMnWgcrnVBtjFdfmzQNtNddjNF
	hFhfPghppPhpRNhzsjsvHVzjpsGnWz
	tTjlCCwMqtdMjMctGJWHwWnVwWnwvWGs
	rZdrjBBtqdCtlcdgFZQLfhRLFSgRNP
	RDHSWrJWffJFlJCgCMCDjCvzjPMP
	QtGTndBwBtNzBVjBCMgB
	LdwwMpTdwsRHsqSHqHJl
	RfsfzvLLFvFzCSvSbDsTpTGMPMZPPTMt
	jqWBjwBBNwWqwPGZbTwVwVtD
	BnhgglhhNNngqjBjHNWrZLlFLSCJSFFCCQzQvQFCFF
	HLvLDQbvnDQDvbHTLhntSnGBSlfGldddcmfMMf
	NgFjZjrZZJrlfJfSVcBJGc
	scWCNFZpsjzrDLwLhbQzhQwD
	SlqJlThDPqpwSTwhcbDdbWDbZGcZNcDb
	MsnWWjHjvLvfscjjgdzNdbgbcc
	vQQvWVQFLLHfHVBWfsfmFFpJRhhSplqlRJqpBwlqTCPC
	DZbDzzZDjQbPGZFFSSgSlFCzTgzm
	qLnvwvhddrqMrwrCTLLFJjmtSlFlSH
	VdhvsWqdVWvvRhsvqbpbPcZfPpjZGBQNRj
	mJNtNFmzDZtzdzrLtwwRqJSchgfGcRfwRB
	pWpjQjCTQnHMWCCpjQpHvTqcwTwScfRcBcSGBRThwS
	MQHjvjVCCqsvljWnVQzLtNPZzmzLVNLddtPN
	QVRPRVDgsRjLssnL
	TTGDJDJfbfLHSnsMWWbs
	qGqqTFFDqgQgQQQq
	nlMnRRjbMjCdJVQJCZ
	nGqfLwfNLFNLnPPGFVVCdVGZJtCtCCVzJz
	LHHfPNHnPqqLwqPqDPWfNFvMglbhhbMgmclgcllDmgmrcl
	cLLWWSThtdLpRcddcgPRZFDMCVPPMCCPCPCZ
	NfGbGNzrBNffGNJjbPPZsZmZZPmDHpMH
	zlJBfzlQzNjNjfJcpwSdvWhcvLwQWt
	cVVQfVCJVrVcTJnfNvlDFmDrmlvrFWlL
	snZHpMhZtMbtPNvzHWWvNFNvNW
	gppnbbbRgMnZbswRqRwbqTcCCSTCJJdGjgfVGTdcCG
	jplgNdrHrrNZgdHmlHNJHddlDSPPSTlzTSlTSDSzCQLfzf
	vscvWWWvGWGGscbFMpRWFwQTPzfLQwQwPfLbzSzzDL
	GvGBWpqcMVRNNZHgdHdtBJ
	LchbZhjjZFjwSmPRqRffqbdtggdR
	vWHMWlHJdGqtRqHV
	MvzCJlnMnlTNnNNLLdhjjCdjjhDjjL
	FNCllHFvCGvwQcPQJfgfmwgh
	zjtRpbDLjtsrzbLLQmfBTgTBQQfhbfQB
	WLgqRzqsrWvFGFZFZC
	qjLlNcLjcNWpQLlQMmvmhCvCgsMZZghj
	tGSDJtRGJzHMMGDVZCfvmfhzmZZgZsmv
	BSSRDRHBGHtSSSbGJSwHbNcLQddqMNlrqcMQMldBWc
	JSfctrtctDpszHvzVQHr
	glCWjhWmFjlmlhmdWPhVVznvcHjszbvvpHvznv
	FgBmFhCBCGFqglgmhCFmSTSRLJLLZfSRJcDSGMtM
	vZGlFFtLMLdShSSShRVtVf
	rQNvmznWPNCPNsrCsbWbsPCvjShhhfHBBHJjSJRhjSRnHhSj
	mCNsQCmqszNcQzrzrrzWvGgGMgpdFpMLlFZGwcLDdg
	QJRJQDlcqLlWbNGL
	HCnwwsCrnstLWqtWNgZNgg
	rsnTrTCHTnnVwnsVPqqDQcRjcczMPvPRzM
	qCzjqnzVdzrdhnhddDbDBMPttcGBDBDPnc
	sZgRQWHgWHHLsgsRRZsJbpJlDcDGNcTDFtGNFFcJNFPBPBTc
	WggbRQSRRgRSsWWmbHqvVffVwhzvCdmfhmdV
	lhqWcNpQGcNmmHmNPWCsQzQsgrQrBMCMbMVM
	wDLFFDJvSFFZRDZSzCrzTzsRgVWbCrMW
	dFwDtZfdjFZWFFfmHGPnPPmqfmPNcN
	lcMRNJRGGLJnNVFbVrwrwZrD
	tjCzQjQhQwgWFShVFS
	ffHQsQssQTzBsPnLpMPRwsJP
	MQSMSBSRFMQLJChLChjTBh
	WmVlPrwnpwDlflNpDrNnDlDwThJCCdLJhhdhCfJTccGjvscd
	gnDVnNnwgglwDwptSZFzgQHqbjZgZZ
	nwBcFgwTDcNrpZMD
	WQWCLZmvhMRvNjsNSD
	CGGWmZGHHhtVzHbTqgTdbgzz
	RmcTCwvssRbsThTcVRJJfSPqfJwJFqfjfMFq
	zQNZDWtQlDZGBQPfFQqjJLjL
	rrglggZGWnrnrrHlDhsbsPTVCsCVsTRpHv
	wFGfzSvCPGttSzqwmtqmvvPRDDRCWgWWDTBTMcBcBWbCRM
	hVJJHQHnpWnDTNnnDb
	LJsVVdhQqvmdbbSf
	srlJztzsVVsSsVtRlNllTWzzmqGhqWLPCDCgmChPLDdqCmCP
	bZQMZpbvMBMgmDGmZLSPZd
	MpScMSMpvfjMBcBcfMfSBnzlTjssNszrNrtlTVzlzFVN
	rCtgrgClprGGClnJCZmwtMjZRjbjjcjZQv
	PWVfBHWPdbNfbbRmRj
	sPsVqFPsHWLhBVVqHFqPVddWSDLJgpTCnnrRRLGpJSSTRrgT
	zjqpGjrQjGqSHCVvCrRZDN
	cTdshMhdmcMNmddRHBhvCCBCCvHZDC
	JTmTmJnLTdwzNQpPWJWgpP
	BmpZmrzZnznHbpprSbQSQbqdSVqbPQcV
	fRGTGJZRTTDwJTJRGDfgJgNFlSSFcldfdccFVlPlFFQPSQ
	GvTTTZZLmsntzmCL
	VhMcrmbhvzMSnhvftbRbllLtglBBtf
	HqqqJqDqPjJPNjjDVFDZCdqBtRtGBGlGRfQQgttQfHlTQl
	pCZJPqqZpmhvhpVh
	dWLBJHJhGJGMBJRcDLDSQsSQpvcR
	ZlnnPqglblfRRpSvSsnz
	sPTgZVjjmwVTljrwTTlbwVGdJhBNNdFdMGNHHJMjBNFN
	FhFrfbfgbLRdfqfrmvDgLdjrcQtSNStHHHQlSjJJPllt
	CnspzZWTpCnMVzzZZGZRCzttHNjNlQlSNtNlNjVcjlQS
	GCZsZBRwnvwfbqwFwb
	bZnJFJgLFRnqQZqJQJFQGpCLNcGlLllClNtccjGc
	rVfvwPDhPHGtlcbClr
	mBhshsfMvBvqsQJdTbgnqQ
	jgWHqMSWMGqWjWjqbWGJQDfVqLfrfDfJhVLfTr
	pPplwsRZPFZFtLhfwgfwrhJL
	zlRsdgFcRgmjdBCMHdjHWB
	qJSGJSPQWzcprtQZtt
	mBMVfsNBnZzcNtcc
	LMLBsmMlvBgFsghVVvfgLBvbJJSqgGHqPGPtCWwbJHqCPG
	ZvZLcdMGVMlHDvDpvqhH
	NNSrQNbJbrTnnWZDDZqqhqpW
	wbgNJrsrCwwJQZbsrJBFzjCCdzGdjcGzMdzj
	JbVmdVLJJJdQMnzmmMgHjPqqjNgvqwngHNNP
	ZfffDZZsRpcpRDcCRrlpplcWSSgwgSwjvvsjPSwhNSWggh
	cCtfppZrpjtMMmdQQTLz
	TtbnmbdmTmgTlPNhqvqj
	wrwrLsVZRsJJJsfHjvPPWfhjHqRN
	sDZwDvsCCQLJZQJQsMCMzZBtSMpndcSFnnSBFtSBmdBc
	mWFTZdmQdZFrFQbCRsrspjSjnvCLRS
	GwlDqcNHDzwGfHSRqCgJsSpnvpSL
	NGlcNwHLLGfDDHDhDwDcwVczbPddZtMFWttWWtdPPdQdhPWd
	mnfcZgcdZqnqdfFqPmHfhqsbgVMCJNMtvCJtMvtblTJtvb
	rRLDDjPSjjPDGBQSBNbtLVtbMNNJlTMtbl
	SzjDDzRRpGQDDDPHzdsmnnhsqcqdFq
	ZDGNRDGjSdwnnmnsVNsHJJ
	tMBWWrddLPLhvWTTPLccvmmbVpgsJHmccppJ
	ClPrtBWWrhrFLBPlCRzjzGqdRzjRdRGZjF
	csTRNQNJcNBDLfhfMf
	qGmWpGHqrqPLChPRhVFPDD
	tgHrtnrrJnZRTZcv
	FLqrfmLDrqCmqjTqcbGqRTGVvb
	FMtWMSWzzFStJzPzhWzhQvTvHVjjTjHTTHvbHc
	PgtWWstWtSpZWPzWwnrBsdBDdFLfllLlfC
	mThbMDMQDCDbwLqWpqPpdhwR
	zgrcffgHNZltZSgHLsRsLLWRWgLqppsW
	SVlSrfSHlSSVlrJfVctlNDMCmMFbnbRDbDBFJFbBRM
	PrBrWqtRPdBLLrBwqpswgpwhgpnZhhzsgw
	FTFRSVJQVJflFfQQgggGMZngGQZszZ
	TbmfFJFSDFblSTDSFFbmVSDrPLLWtcmBqqRmBtmcLtcrjP
	DjPsMwDjLVVTsvNNRTNTRT
	ztdQQHqHlFNtfRNNNMgg
	FzhMhHQlDcCrhCCc
	zSHGzzmHgnnMDLTNTG
	lPVBtvhQjpNSMWTLBD
	VCftbjvbVCfPbZwsJsrSgSSZwC
	CbwgmvMnmnCwMmwRQqJBGBgHZHpJHdtdZpJt
	zVSlNSDlrzNhqlNTScDzVWfBBZZZZGBstGsdsWFpdHdJsW
	NDlLzhrVcqRPCMRwLLLw
	TjTHHLwnLjVlTwLjgVfvsFvDsdWfvDvFMd
	qbRRRpmpcmDcczppztSqSvWFssFGfWdMvfQWdfsG
	RZpqDBmtrzhzphjTgjHlnwjgJhgJ
	dLmMgdgzwDLzDWFhBWvzFzzBZJ
	tTVcppbSTfstTMMHfTbhBchhJFCWcjWBZhjGGB
	SSSSNbsNRpRRsRrfVHfRpNtlPgQDLPdMmlDLlrPnqPdPLl
	qqbTCSqdqqFZdRLZhwhZ
	HWWlHtlrBfGtVssnsLnHfJVPPMMFzhPRwMPwFhzPZzPMGM
	nfmtsrlsnrfVnHJrVBWlsVfgbbNTNSvmvvpcTjLjLbqvvS
	GGhFvGPFcThqffPdnfNLqZZCSwtQSwZpwQQBsL
	RglMRrJJgHBCBZSQQpdr
	WmbRHHbzDgJMDzRDMdWmWHzHNFFvvGGhnvVvvfcvnFfcbvnT
	QsfQmsLfZZZcshnJ
	dSgdWgSVVFvzSpqFdqTgWRHbJNcbZNCTJCNNZRRCCh
	FcpVjgDvVVFdVWFvzjwwQtBMLtBBGDwftPrB
	rqsRrHsvsPqswNcJcNJrnnBrNn
	bFjgGFdbVRNNnpRQpV
	GSthhggGDSvMRqtHvMfM
	ZwVPgMsgVsGzVsRZpgpzzgpFMrNbbLFrDLFFrrSDLfrNBN
	qvnjBhQhntbfDLrF
	CJlHHcHcTWqvpBdsWRpdPdgs
	BjmTDjJBCBWrgQRPFlWWlW
	dHphshtdtVHVhpJqspdvRrqFPgrLPPFPrrRPvQ
	sdMsMtStVszpwMzHjJGjCcZjmScNfCDf
	DmGdDffgDSDDdJstqdJldlRt
	MhnvMCZCbbZHMvsCHtrcVrPjJcRqVtlt
	LsQbsFZvZhQzZwhQWTNgBWpNwSGpTmfS
	RRJQnCzbZZLTZJCBtWvFtsfqBqtfWb
	prjlChGNldGNdlSVMhWfqWtfsvwvqsFtdtsq
	GGjNDNhpMGMGVhrnZZTzcTHCCJcDHc
	RmbMmjgpPjMBsBMfchhVsc
	HwFWFTztSrtFpcQvBsSqVscBBC
	zWwnJFHtWWHDgbGgdpGpnl
	mnbWbRRLRFnmmWcCDTBVwCDBlwNW
	ggJPtpdHGfdZtMHgtZgVPPBCVsPNBcsBTTDDCC
	hpvJJTpGhdhtJdMHqvmmnLvSbmnFnRFm
	WWtrWrNgVbRjMrQCNzqJFwQJFNTJ
	LdHPhcdchQQssLzJrz
	pBccnHpnrrcGHnnSlWjnRMSlbt
	NMMfNFnZgMVThhTMcgTDJDJjsVvvJJqJmHsqHG
	LQpwwprCQzBNBdGjGjHswswdvm
	CBCzzCrbWbSlNQnTRgPPfFRWnfgc
	RFwHVQRwFgTQSFVhdsdHsBdDBnnqnq
	LGftLtPGGMzlNrhlPqPsrJ
	fvGpWpMtccpTwwpRRQhh
	TTJCGdTGtZRQQCnzcnCv
	FWWHPSFNFbDbDDqSWnVmLRRjRRQLhcmLjS
	qPwPWwFppbwggGZGfdJZgdnGdd
	zSTWzrzWTLWpCtCGpqqGgplc
	nZWwsJVZZBnJHJCclHllgtChgCgc
	DFnVBJsFssVVFBFnBdfvjDSmTMWzrmMfRmTv
	MJmgMssrsggqqMVstbwTcTbPbTTwThmw
	NRBBGRjHVRRcRbCp
	QnSfzLWzNHzNVQQVjrglJMsMFvgJdFWrgZ
	ggLLGnhgnPvJHZnN
	VBtmVSldbSBVlcNPHvjmNcwNZZ
	tdWqSVSSBztVWGrThLhfrfvG
	TDqrjdSwLqDppdTCdzPBFmmjQmhHFPFQhPFR
	zlGbMcVcVtsPHFRhWRRsPF
	btgvlVVcDZZZqgrz
	DgwlgbbFDDjjPTHDrmddPhPV
	WqtMBBtQsttMNWQBqsbJpGGzdPdTHLVmTzJhmTPhHHPTmH
	qQsqGZNQtZGMNsNtZpFnjnCRbZffwwSRljFf
	gMdFLCdnMZCTFFCqnTgWLCHfSgPgPHStcQQmfSBBSfHg
	vrwwrwzbGjjswjvhGGsjPQmqRmHPbBtcBQtqfmcH
	qzJllVsGVGljjsrzwDzhwzDGTddNLFnZWNdpCVWTNTZTLZCF
	LtwMhDtctwbwwppdWBJQJBWPvPfDfqvG
	FTzrNrgSRFrgzFRHNVFQJvlqHjBvQWlQWqPBfq
	sFgNzmVmNzgTvVTMwhMhstMwZtsbsc
	MrBDQVzzlrvhQzQrDMVQrzrzgRJnRRwwRbwSwwVRRNSgwwwJ
	qFTPTvfTHcqqncpcwR
	LmtdGGPmTPGCTLHLWsZMhvZMMMzrzzdlMQ
	ZVNpjfpZNpfNgNjzNVfWtnbbWmBHtsZWBSZBGS
	MrDrQvvDrPLDMvFvdmBGGsBBCtsHrnrGCm
	ltRMwLLDDRlvQwvlQcwhqfcJNpgzjJpjhJ
	sRRRlRbcFbBBdnFBwCGppNvGrTCDDGVNlr
	PPSLQzHjzZZPLZPjgTNTgpCbVJvGrNCTGr
	ZLHHPQjhQmWWSRRnssdtbnmfwF
	GRwrMrHJGwJPGWsgfqQgsc
	VbTvLQCZLSWWsgWf
	TVDvVCvppvTDmzZVTbZpTzBBNQQQJlJBBJBNNJmRBwRH
	shJRWJsjZGNjSTrjFS
	dMLCddggldQzMCCVgzVVLmLvTwNFFSqpNSqSbFGSqTTpMTFN
	VGQvVglCLcVzgdddCDVvlsPZRRBDJPHZWZZnBsWJRR
	CrwlwhRCMrswnsHBFccHHWFc
	QJTmtfQgLtzQfLQfdPcWSFHHDDSpcFpFBg
	jTQTqbfQfmLbLQJbJrRCWjljZGjNrZlZlC
	JmthDmLShtJmHphphJQCwjdjdFDzFgzFdgdNlC
	sbMTVBrWMbNvVMnsWMnVzjsjwCfjFgfZzfdgdzlj
	NvqbbBcMMPPSqLSpGGthmp
	RfGWFHlPFFNWGFZRZBjvwCvzBwhhrvvjzmrr
	sLJSLMSTSJTbStJtMSqSqbpMrvmrzWdvhmjDCzzwrrpjdDDv
	SbQqsqsWcZPcQGFG
	BjqbMqMVBsfqGqFqGLmF
	ZZQbQPddPcwbPnRQltdtQZdnmFNrvfhGrhrWWFNWWtmNFNNW
	dJJQccnRPpcbQcMHsSgSMsDMTJSg
	WWGBBvPflnWbBWhvhbPvNfnnVCFZmVRVZmVGMVwRLCCCGwVC
	gjszgTMrgzgqCRRdmJRjJLVw
	grzQHzqczMSzqSHcgQsqPvPlbNblpPhhPPbHvnhp
	sJDDNWdnRLTTvqwSFPCmLCCrCq
	thzplgfjglflFcbMclpppMfcwPqCZQCmqCwrzCqmQmHSqPqq
	MhcpFBMBlhjbBTdnNJWvNvsvBd
	czwwghnWWfcfgwfWthfrvVvrjdrdvDDVrbzrLF
	RHPPMRpQPRMPPJRjJQsZsrrvvJBDDVDVdFqrBrFdBv
	smjMsGZHRsHSmRQNGHPpSTwwttCflwngnChcCtWW
	bprrrwrtLDtrWwrQjRDQDbPPVHVmmmmHNWlPlVNPZZlv
	hqqhfnBCTfnnhzJwzsqzfPZZMCCVZVHHFvZMFvZmlC
	TzhhdJTqJzcBdJJnzjtQrLdjwgLtpbgrLQ
	qzQvzzgWSCqtqqGpddGc
	jLrZNZhZrNRLHNffhrjNjNdtdZtGcPFwFwpbGwbVpdwC
	nHnhrLNCCMHmhHBMhrzvgJvsWSWMWzzWzSlv
	RzcbzdRFzbbzbzbFdZFTHMZPhVhVQMLrlrQPhLZlMM
	BNGfBvsNttVmMhlMLm
	BwGjpllswfjwpcFDWcWcbpdb
	SjzpswrLSDjVSpwlmZJBTBdNJLvBNvHQZT
	rCcCtbqgCfthggtbGGMqqghqZQvvQTBNJQHQZQTcZTJFZFFd
	CggGMtqMfWbbGghPhhbCMtmsSppSspjpmWzjVSWlVrrm
	PmWTPThTQWnLWQFl
	VNcSVfMbtsddBQNnNpdl
	sSjctwjVSzzccjgnTnDTHRDhqjRR
	WfMWfCNCjWWHNTccMjRjfRcMbqSwfVwqwsfGGbssrJSrswVw
	llLFQLlvlPFnhQBPBZQBqvBwzSzGGhShJVwShmsJbbmzSG
	lnPqvQZBFFBnnpgplFvtvHDjTdcTjTMMjCRNCMWgRC
	rprFNFFNjNLmMdgcqL
	BvzCQQbBQgffsDbvVHMdbcVqmLVqlmqq
	JvJCzBDJwnsRnQDszCBnnnQBrjZPjFpgZFTFZRpTrpZFGFtT
	wBHQQZHVCcpwDgdZdMsZjvMZFn
	GPSzlNlJLfzzzvsWdWLMmFWLMM
	NfqGSfrTNzRTqJfRbptQHFQFrwrFHBHw
	sNjVMVNVMzPzQgghcMsNzJtjSJtTFDTJtJnnDLjDnL
	CHwrdCpvCrwrWdpZqcpFttJSFJTLLHLJfbnbfD
	qrlZCwlqZrqqpWdlRqCRqdqcVNsVMzQzmNgNPBsRhVQVVzMs`
}

// part 1
// func getCharCode(s string) rune {
// 	return []rune(s)[0]
// }

// func findDuplicate(rucksack string) int {
// 	midPoint := int(math.Floor(float64(len(rucksack) / 2)))
// 	comp1 := rucksack[0:midPoint]
// 	comp2 := rucksack[midPoint:]

// 	duplicate := []string{}

// 	for i := 0; i < len(comp1); i++ {
// 		for j := 0; j < len(comp2); j++ {
// 			if comp1[i] == comp2[j] && len(duplicate) < 1 {
// 				duplicate = append(duplicate, string(comp1[i]))
// 			}
// 		}
// 	}

// 	if len(duplicate) > 0 {

// 		// if char is lower case
// 		if strings.ToLower(duplicate[0]) == duplicate[0] {
// 			return int(getCharCode(duplicate[0])) - 96
// 		}
// 		// if char is upper case
// 		if strings.ToUpper(duplicate[0]) == duplicate[0] {
// 			return int(getCharCode(duplicate[0])) - 38
// 		}
// 	}
// 	return 0
// }

// func main() {
// 	rucksacks := strings.Split(getInput(), "\n")
// 	total := 0

// 	for _, rucksack := range rucksacks {
// 		total += findDuplicate(rucksack)
// 	}
// 	fmt.Println(total)
// }

// part 2
// pass three strings at a time
func getCharCode(s string) rune {
	return []rune(s)[0]
}

func getLongestString(p1 string, p2 string, p3 string) (string, string, string) {
	if len(p1) >= len(p2) && len(p1) >= len(p3) {
		return p1, p2, p3
	}
	if len(p2) >= len(p1) && len(p2) >= len(p3) {
		return p2, p1, p3
	}
	return p3, p1, p2
}

func findDuplicate(p1 string, p2 string, p3 string) int {
	dups := []string{}
	longest, s2, s3 := getLongestString(p1, p2, p3)

	for i := 0; i < len(longest); i++ {
		if strings.Contains(s2, string(longest[i])) && strings.Contains(s3, string(longest[i])) {
			dups = append([]string{string(longest[i])}, dups...)
		}
	}

	return getAsciiVal(dups)
}

func getAsciiVal(dups []string) int {
	if len(dups) > 0 {

		// if char is lower case
		if strings.ToLower(string(dups[0])) == dups[0] {
			return int(getCharCode(dups[0])) - 96
		}
		// if char is upper case
		if strings.ToUpper(dups[0]) == dups[0] {
			return int(getCharCode(dups[0])) - 38
		}
	}
	return 0
}

func main() {
	rucksacks := strings.Split(getInput(), "\n")
	total := 0

	for i := 1; i < len(rucksacks); i = i + 3 {
		total += findDuplicate(rucksacks[i-1], rucksacks[i], rucksacks[i+1])
	}
	fmt.Println(total)
}
