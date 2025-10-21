select c.iban, b.amount, COUNT(t.customer_id) as transaction_count
from customers c
join transactions t on t.customer_id = c.id
join balances b on b.customer_id = c.id
where b.amount < 0
group by c.iban, b.amount
order by b.amount asc;
