; ({col4} < 0) and ({col3} =~ "支付宝")
2022-{{substr 0 2 .col2}}-{{substr 3 5 .col2}} *  "{{.col3}}"
  Expenses:Live {{.col4}} CNY
  Liabilities:CreditCard:CMB {{mul .col4 -1}} CNY