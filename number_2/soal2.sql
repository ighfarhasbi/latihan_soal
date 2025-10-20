select 
fullname || ' ID: ' || employeeid || ' has a performence rating of ' ||
case 
	when performencescore < 50 then 'Need Improvement'
	when performencescore < 75 then 'Meets Expextations'
	when performencescore < 90 then 'Exeeds Expectetions'
	when performencescore <= 100 then 'Outstanding'
end as result
from employees
order by employeeid asc