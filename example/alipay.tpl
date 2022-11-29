; {col1} == "收入" and ({col6} != 0)
{{toDate "2006/1/2 15:04" .col11 | date "2006-01-02"}} * "{{.col2}}" "{{.col4}}"
    Assets:Alipay {{.col6}} CNY
    Equity:Opening-Balances

; ({col1} == "支出") and ({col8} =~ "物业") and ({col5} =~ "花呗") and ({col7}=~ "成功") and ({col6} != 0)
{{toDate "2006/1/2 15:04" .col11 | date "2006-01-02"}} * "{{.col2}}" "{{.col4}}"
    Liabilities:HuaBei -{{.col6}} CNY
    Expenses:Home

; ({col1} == "支出") and ({col8} =~ "物业") and ({col5} =~ "招商银行信用卡") and ({col7}==~ "成功") and ({col6} != 0)
{{toDate "2006/1/2 15:04" .col11 | date "2006-01-02"}} * "{{.col2}}" "{{.col4}}"
    Liabilities:CreditCard:CMB -{{.col6}} CNY
    Expenses:Home

; ({col1} == "支出") and ({col8} =~ "交通出行") and ({col5} =~ "招商银行信用卡") and ({col7}=~ "成功") and ({col6} != 0)
{{toDate "2006/1/2 15:04" .col11 | date "2006-01-02"}} * "{{.col2}}" "{{.col4}}" #walk
    Liabilities:CreditCard:CMB -{{.col6}} CNY
    Expenses:Walk

; ({col1} == "支出")  and ({col8} =~ "交通出行") and ({col5} =~ "中信银行信用卡") and ({col7}=~ "成功")  and ({col6} != 0)
{{toDate "2006/1/2 15:04" .col11 | date "2006-01-02"}} * "{{.col2}}" "{{.col4}}" #walk
    Liabilities:CreditCard:CITIC -{{.col6}} CNY
    Expenses:Walk

; ({col1} == "支出")  and ({col8} =~ "交通出行") and ({col5} == "花呗") and ({col7}=~ "成功") and ({col6} != 0)
{{toDate "2006/1/2 15:04" .col11 | date "2006-01-02"}} * "{{.col2}}" "{{.col4}}" #walk
    Liabilities:HuaBei -{{.col6}} CNY
    Expenses:Walk

; ({col1} == "支出")  and ({col8} =~ "交通出行") and ({col5} == "余额宝") and ({col7}=~ "成功") and ({col6} != 0)
{{toDate "2006/1/2 15:04" .col11 | date "2006-01-02"}} * "{{.col2}}" "{{.col4}}" #walk
    Assets:Alipay -{{.col6}} CNY
    Expenses:Walk

; ({col1} == "支出") and ({col5} =~ "招商银行信用卡") and ({col7}=~ "成功") and ({col6} != 0)
{{toDate "2006/1/2 15:04" .col11 | date "2006-01-02"}} * "{{.col2}}" "{{.col4}}"
    Liabilities:CreditCard:CMB -{{.col6}} CNY
    Expenses:Others

; ({col1} == "支出") and ({col5} =~ "中信银行信用卡") and ({col7}=~ "成功") and ({col6} != 0)
{{toDate "2006/1/2 15:04" .col11 | date "2006-01-02"}} * "{{.col2}}" "{{.col4}}"
    Liabilities:CreditCard:CITIC -{{.col6}} CNY
    Expenses:Others

; ({col1} == "支出") and ({col5} =~ "花呗") and ({col7} =~ "成功") and ({col6} != 0)
{{toDate "2006/1/2 15:04" .col11 | date "2006-01-02"}} * "{{.col2}}" "{{.col4}}"
    Liabilities:HuaBei -{{.col6}} CNY
    Expenses:Others

; ({col1} == "支出") and ({col5} in ["余额宝","余额"]) and ({col7} =~ "成功") and ({col6} != 0)
{{toDate "2006/1/2 15:04" .col11 | date "2006-01-02"}} * "{{.col2}}" "{{.col4}}"
    Assets:Alipay -{{.col6}} CNY
    Expenses:Others

; ({col1} == "支出") and ({col5} =~ "北京银行储蓄卡") and ({col7} =~ "成功") and ({col6} != 0)
{{toDate "2006/1/2 15:04" .col11 | date "2006-01-02"}} * "{{.col2}}" "{{.col4}}" #cat
    Assets:Card:BeiJing -{{.col6}} CNY
    Expenses:Cat


; ({col1} == "其他") and ({col8} =~ "物业") and ({col5} =~ "花呗") and ({col7}=="退款成功") and ({col6} != 0)
{{toDate "2006/1/2 15:04" .col11 | date "2006-01-02"}} * "{{.col2}}" "{{.col4}}"
    Liabilities:HuaBei {{.col6}} CNY
    Expenses:Home

; ({col1} == "其他") and ({col8} =~ "物业") and ({col5} =~ "招商银行信用卡") and ({col7}=="退款成功") and ({col6} != 0)
{{toDate "2006/1/2 15:04" .col11 | date "2006-01-02"}} * "{{.col2}}" "{{.col4}}"
    Liabilities:CreditCard:CMB {{.col6}} CNY
    Expenses:Home

; ({col1} == "其他") and ({col8} =~ "交通出行") and ({col5} =~ "招商银行信用卡") and ({col7}=="退款成功") and ({col6} != 0)
{{toDate "2006/1/2 15:04" .col11 | date "2006-01-02"}} * "{{.col2}}" "{{.col4}}" #walk
    Liabilities:CreditCard:CMB {{.col6}} CNY
    Expenses:Walk

; ({col1} == "其他")  and ({col8} =~ "交通出行") and ({col5} =~ "中信银行信用卡") and ({col7}=="退款成功") and ({col6} != 0)
{{toDate "2006/1/2 15:04" .col11 | date "2006-01-02"}} * "{{.col2}}" "{{.col4}}" #walk
    Liabilities:CreditCard:CITIC {{.col6}} CNY
    Expenses:Walk

; ({col1} == "其他")  and ({col8} =~ "交通出行") and ({col5} == "花呗") and ({col7}=="退款成功") and ({col6} != 0)
{{toDate "2006/1/2 15:04" .col11 | date "2006-01-02"}} * "{{.col2}}" "{{.col4}}" #walk
    Liabilities:HuaBei {{.col6}} CNY
    Expenses:Walk

; ({col1} == "其他")  and ({col8} =~ "交通出行") and ({col5} == "余额宝") and ({col7}=="退款成功") and ({col6} != 0)
{{toDate "2006/1/2 15:04" .col11 | date "2006-01-02"}} * "{{.col2}}" "{{.col4}}" #walk
    Assets:Alipay {{.col6}} CNY
    Expenses:Walk

; ({col1} == "其他") and ({col5} =~ "招商银行信用卡") and ({col7}=="退款成功") and ({col6} != 0)
{{toDate "2006/1/2 15:04" .col11 | date "2006-01-02"}} * "{{.col2}}" "{{.col4}}"
    Liabilities:CreditCard:CMB {{.col6}} CNY
    Expenses:Others

; ({col1} == "其他") and ({col5} =~ "中信银行信用卡") and ({col7}=="退款成功") and ({col6} != 0)
{{toDate "2006/1/2 15:04" .col11 | date "2006-01-02"}} * "{{.col2}}" "{{.col4}}"
    Liabilities:CreditCard:CITIC {{.col6}} CNY
    Expenses:Others

; ({col1} == "其他") and ({col5} == "花呗") and ({col7}=="退款成功") and ({col6} != 0)
{{toDate "2006/1/2 15:04" .col11 | date "2006-01-02"}} * "{{.col2}}" "{{.col4}}"
    Liabilities:HuaBei {{.col6}} CNY
    Expenses:Others

; ({col1} == "其他") and ({col5} == "余额宝") and ({col7}=="退款成功") and ({col6} != 0)
{{toDate "2006/1/2 15:04" .col11 | date "2006-01-02"}} * "{{.col2}}" "{{.col4}}"
    Assets:Alipay {{.col6}} CNY
    Expenses:Others

; ({col1} == "其他") and ({col4} =~ "收益发放") and ({col7}=="交易成功") and ({col6} != 0)
{{toDate "2006/1/2 15:04" .col11 | date "2006-01-02"}} * "{{.col2}}" "{{.col4}}"
    Assets:Alipay {{.col6}} CNY
    Income:Interest

; ({col1} == "其他") and ({col4} =~ "单次转入") and ({col5} =~ "北京银行储蓄卡") and ({col7}=="交易成功") and ({col6} != 0)
{{toDate "2006/1/2 15:04" .col11 | date "2006-01-02"}} * "{{.col2}}" "{{.col4}}"
    Assets:Card:BeiJing -{{.col6}} CNY
    Assets:Alipay

; ({col1} == "其他") and ({col4} =~ "单次转入") and ({col5} =~ "招商银行储蓄卡") and ({col7}=="交易成功") and ({col6} != 0)
{{toDate "2006/1/2 15:04" .col11 | date "2006-01-02"}} * "{{.col2}}" "{{.col4}}"
    Assets:Card:CMB -{{.col6}} CNY
    Assets:Alipay

; ({col1} == "其他") and ({col2} =~ "招商银行") and ({col5} == "余额宝") and ({col7}=="还款成功") and ({col6} != 0)
{{toDate "2006/1/2 15:04" .col11 | date "2006-01-02"}} * "{{.col2}}" "{{.col4}}" #payback
    Liabilities:CreditCard:CMB  {{.col6}} CNY
    Assets:Alipay

; ({col1} == "其他") and ({col2} =~ "花呗") and ({col5} == "余额宝") and ({col7}=="还款成功") and ({col6} != 0)
{{toDate "2006/1/2 15:04" .col11 | date "2006-01-02"}} * "{{.col2}}" "{{.col4}}" #payback
    Liabilities:HuaBei  {{.col6}} CNY
    Assets:Alipay

; ({col1} == "其他") and ({col2} =~ "中信银行") and ({col5} in ["余额宝", "余额"]) and ({col7}=="还款成功") and ({col6} != 0)
{{toDate "2006/1/2 15:04" .col11 | date "2006-01-02"}} * "{{.col2}}" "{{.col4}}" #payback
    Liabilities:CreditCard:CITIC  {{.col6}} CNY
    Assets:Alipay

; ({col1} == "其他") and ({col2} =~ "花呗") and ({col5} =~ "招商银行储蓄卡") and ({col7}=="还款成功") and ({col6} != 0)
{{toDate "2006/1/2 15:04" .col11 | date "2006-01-02"}} * "{{.col2}}" "{{.col4}}" #payback
    Liabilities:HuaBei  {{.col6}} CNY
    Assets:Card:CMB

; ({col1} == "其他") and ({col2} =~ "西安银行") and ({col4} =~ "转出到银行卡") and ({col7}=="交易成功") and ({col6} != 0)
{{toDate "2006/1/2 15:04" .col11 | date "2006-01-02"}} * "{{.col2}}" "{{.col4}}" #walk
    Expenses:Walk  {{.col6}} CNY
    Assets:Alipay

; ({col1} == "其他") and ({col2} =~ "建设银行") and ({col4} =~ "转出到银行卡") and ({col7}=="交易成功") and ({col6} != 0)
{{toDate "2006/1/2 15:04" .col11 | date "2006-01-02"}} * "{{.col2}}" "{{.col4}}" 
    Assets:Card:CCB  {{.col6}} CNY
    Assets:Alipay

; ({col1} == "其他") and ({col2} =~ "招商银行") and ({col4} =~ "转出到银行卡") and ({col7}=="交易成功") and ({col6} != 0)
{{toDate "2006/1/2 15:04" .col11 | date "2006-01-02"}} * "{{.col2}}" "{{.col4}}" 
    Assets:Card:CMB  {{.col6}} CNY
    Assets:Alipay

; ({col1} == "其他") and ({col4} =~ "结息") and ({col7}=="交易成功") and ({col6} != 0)
{{toDate "2006/1/2 15:04" .col11 | date "2006-01-02"}} * "{{.col2}}" "{{.col4}}"
    Income:Interest  -{{.col6}} CNY
    Assets:Alipay