create table `bill` (
    `id` int unsigned auto_increment primary key,
    `timestamp` timestamp not null,
    `in_ex` varchar(6) check(in_ex in("收入", "支出")),
    `commodity` varchar(150),
    `commodity_type` varchar(60),
    `ammount` decimal(10, 2) not null,
    `pay_method` varchar(60),
    `counterparty` varchar(60),
    `tran_account` varchar(60),
    `order_number` varchar(30),
    `mer_number` varchar(30),
    `source` varchar(30),
    `remark` varchar(255)
);