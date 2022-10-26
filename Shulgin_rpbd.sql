/*1.Выведите на экран любое сообщение.*/
DO
$$
BEGIN
    RAISE NOTICE 'Message';
END
$$
/*2.Выведите на экран текущую дату.*/
DO
$$
BEGIN
    RAISE NOTICE 'Curent date: %', CURRENT_DATE;
END
$$
/*3.Создайте две числовые переменные и присвойте им значение. Выполните математические действия с этими числами и выведите результат на экран.*/
CREATE OR REPLACE PROCEDURE sum_two(INOUT x int, y int)
LANGUAGE plpgsql
AS $$
BEGIN x := x + y;
END;
$$;

DO $$
DECLARE a int := 2; b int := 3;
BEGIN 
    CALL sum_two(a, b);
	RAISE NOTICE 'sum = %', a;
END;
$$;
/*4.Написать программу двумя способами 1 - использование IF, 2 - использование CASE. Объявите числовую переменную и присвоейте ей значение. Если число равно 5 - выведите на экран "Отлично". 4 - "Хорошо". 3 - Удовлетворительно". 2 - "Неуд". В остальных случаях выведите на экран сообщение, что введённая оценка не верна.*/
/*IF*/
CREATE OR REPLACE PROCEDURE def(INOUT x int, INOUT y varchar)
LANGUAGE plpgsql
AS $$
BEGIN 
    IF x = 5 THEN
        y := 'Отлично';
    ELSEIF x = 4 THEN
        y := 'Хорошо';
    ELSEIF x = 3 THEN
        y := 'Удовлетворительно';
    ELSEIF x = 2 THEN
        y := 'Неуд';
    ELSE
        y := 'Некорректная оценка';
    END IF;
END;
$$;

DO $$
DECLARE a int := 4; b varchar := '';
BEGIN 
    CALL def(a, b);
	RAISE NOTICE 'Оценка = %', b;
END;
$$;
/*CASE*/
CREATE OR REPLACE PROCEDURE def_case(INOUT x int, INOUT y varchar)
LANGUAGE plpgsql
AS $$
BEGIN 
    CASE 
        WHEN x = 5 THEN y := 'Отлично';
        WHEN x = 4 THEN y := 'Хорошо';
        WHEN x = 3 THEN y := 'Удовлетворительно';
        WHEN x = 2 THEN y := 'Неуд';
    ELSE
        y := 'Некорректная оценка';
    END CASE;
END;
$$;

DO $$
DECLARE a int := 4; b varchar := '';
BEGIN 
    CALL def_case(a, b);
	RAISE NOTICE 'Оценка = %', b;
END;
$$;
/*5.Выведите все квадраты чисел от 20 до 30 3-мя разными способами (LOOP, WHILE, FOR).*/
/*LOOP, FOR.*/
CREATE OR REPLACE PROCEDURE pow_for()
LANGUAGE plpgsql
AS $$
BEGIN 
    FOR i IN 20..30 LOOP
        RAISE NOTICE '%', i^2;
    END LOOP;
END;
$$;

CALL pow_for();
/*WHILE.*/
CREATE OR REPLACE PROCEDURE pow_while(i int)
LANGUAGE plpgsql
AS $$
BEGIN 
    WHILE i < 31 LOOP
        RAISE NOTICE '%', i^2;
        i := i + 1;
    END LOOP;
END;
$$;

CALL pow_while(20);
/*6.Последовательность Коллатца. Берётся любое натуральное число. Если чётное - делим его на 2, если нечётное, то умножаем его на 3 и прибавляем 1. Такие действия выполняются до тех пор, пока не будет получена единица. Гипотеза заключается в том, что какое бы начальное число n не было выбрано, всегда получится 1 на каком-то шаге. Задания: написать функцию, входной параметр - начальное число, на выходе - количество чисел, пока не получим 1; написать процедуру, которая выводит все числа последовательности. Входной параметр - начальное число.*/
/*FUNCTION.*/
CREATE OR REPLACE FUNCTION kolf(i int) RETURNS int
AS $$
DECLARE
    kol int := 0;
BEGIN
	WHILE i != 1 LOOP
    	IF mod(i, 2) = 0 THEN
        	i := i / 2;
        	kol := kol + 1;
    	ELSEIF mod(i,2) = 1 THEN
        	i := i * 3 + 1;
        	kol := kol + 1;
    	END IF;
	END LOOP;
    RETURN kol;
END
$$ LANGUAGE plpgsql;

SELECT kolf(20)
/*PROCEDURE.*/
CREATE OR REPLACE PROCEDURE kol(i int)
LANGUAGE plpgsql
AS $$
BEGIN 
    WHILE i != 1 LOOP
        RAISE NOTICE '%', i;
        IF mod(i, 2) = 0 THEN
            i := i / 2;
        ELSEIF mod(i,2) = 1 THEN
            i := i * 3 + 1;
        END IF;
    END LOOP;
END;
$$;

CALL kol(20);

/*7.Числа Люка. Объявляем и присваиваем значение переменной - количество числе Люка. Вывести на экран последовательность чисел. Где L0 = 2, L1 = 1 ; Ln=Ln-1 + Ln-2 (сумма двух предыдущих чисел). Задания: написать фунцию, входной параметр - количество чисел, на выходе - последнее число (Например: входной 5, 2 1 3 4 7 - на выходе число 7); написать процедуру, которая выводит все числа последовательности. Входной параметр - количество чисел.*/
/*FUNCTION.*/
CREATE OR REPLACE FUNCTION lukef(i int) RETURNS int
AS $$
DECLARE
    L1 int := 2;
    L2 int := 1;
    Ln int := 0;
BEGIN
	FOR i IN 1..i-2 LOOP
        Ln := L1 + L2;
        L1 := L2;
        L2 := Ln;
	END LOOP;
    RETURN Ln;
END
$$ LANGUAGE plpgsql;

SELECT lukef(5)
/*PROCEDURE.*/
CREATE OR REPLACE PROCEDURE lukep(i int)
LANGUAGE plpgsql
AS $$
DECLARE
    L1 int := 2;
    L2 int := 1;
    Ln int := 0;
BEGIN 
    RAISE NOTICE '%', L1;
    RAISE NOTICE '%', L2;
    FOR i IN 1..i-2 LOOP
        Ln := L1 + L2;
        RAISE NOTICE '%', Ln;
        L1 := L2;
        L2 := Ln;
    END LOOP;
END;
$$;

CALL lukep(5);

/*8.Напишите функцию, которая возвращает количество человек родившихся в заданном году.*/
CREATE OR REPLACE FUNCTION birth_year(god int) RETURNS int
AS $$
DECLARE
    cur CURSOR (input integer) FOR SELECT * FROM people WHERE EXTRACT(year FROM people.birth_date) = birth_year.god;
    p people%ROWTYPE;
    k int := 0;
BEGIN
    OPEN cur(5);
   loop
   		FETCH cur INTO p;
   		exit when not found;
   		k := k + 1;
   end loop;
   CLOSE cur;
   RETURN k;
END
$$ LANGUAGE plpgsql;

SELECT birth_year(1995)

/*9.Напишите функцию, которая возвращает количество человек с заданным цветом глаз..*/
CREATE OR REPLACE FUNCTION colour(cvet varchar) RETURNS int
AS $$
DECLARE
    cur CURSOR (input integer) FOR SELECT * FROM people WHERE eyes = colour.cvet;
	p people%ROWTYPE;
	k int := 0;
BEGIN
   OPEN cur(5);
   loop
   		FETCH cur INTO p;
   		exit when not found;
   		k := k + 1;
   end loop;
   CLOSE cur;
   RETURN k;
END
$$ LANGUAGE plpgsql;

SELECT colour('brown')
/*10.Напишите функцию, которая возвращает ID самого молодого человека в таблице.*/
CREATE OR REPLACE FUNCTION get_junior_user_id(weight real) RETURNS integer
AS $$
DECLARE
    answer integer;
BEGIN
    SELECT id INTO STRICT answer FROM people WHERE birth_date = (SELECT MAX(birth_date) FROM people);
    RETURN answer;
END
$$ LANGUAGE plpgsql;

SELECT get_junior_user_id();




