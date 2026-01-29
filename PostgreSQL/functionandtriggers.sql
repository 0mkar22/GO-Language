CREATE OR REPLACE FUNCTION fn_mid2(buffer varchar, startpos integer, len integer)
RETURNS varchar 
AS 
$$
BEGIN
    RETURN SUBSTRING(buffer FROM startpos FOR len);
END;
$$ 
LANGUAGE plpgsql;

SELECT * FROM fm_mimd2('Hello, World!', 8, 5); 


CREATE OR REPLACE FUNCTION fnMakeFull(fisrt_name varchar, last_name varchar)
RETURNS varchar 
AS 
$$
BEGIN
    if first_name is NULL AND last_name is NULL THEN
        RETURN NULL;
    ELSIF first_name is NULL AND last_name is NOT NULL THEN
        RETURN initCap(last_name);
    ELSIF last_name is NULL AND first_name is NOT NULL THEN 
        RETURN initCap(first_name);
    ELSE
        RETURN concat(initCap((first_name)), ' ', initCap((last_name)));
    END IF;

END;
$$ 
LANGUAGE plpgsql;

SELECT * FROM fnMakeFull('Omkar', 'Bhogte');

CREATE OR REPLACE FUNCTION fnswap(INTO num1, INTO num2)
AS
$$
BEGIN
    SELECT num1, num2 INTO num2, num1;
END;
$$
LANGUAGE plpgsql;

SELECT * FROM fnswap(10, 20);



CREATE OR REPLACE FUNCTION fnMean(NUMERIC[])
RETURNS NUMERIC
AS
$$
DECLARE
    total NUMERIC := 0;
    cat INT := 0;
    n_array ALIAS FOR $1;
BEGIN
    FOREACH i IN ARRAY n_array
    LOOP
        total := total + i;
        cat := cat + 1;
    END LOOP;
    RETURN total / cat;
END;
$$
LANGUAGE plpgsql;

SELECT * FROM fnMEAN(ARRAY[10, 20, 30, 40, 50]);


CREATE OR REPLACE FUNCTION fnGetYearofCar(yr INTEGER)
RETURNS TABLE(car_uid UUID, make VARCHAR, model VARCHAR, price NUMERIC, year INTEGER)
AS
$$
BEGIN
    RETURN QUERY
    SELECT c.car_uid, c.make, c.model, c.price, c.year FROM car c 
    WHERE c.year = yr;
END;
$$
LANGUAGE plpgsql;

SLECT * FROM fnGetYearofCar(2020);


CREATE TRIGGER trg_before_insert_person
BEFORE INSERT ON person
FOR EACH ROW
EXECUTE FUNCTION fn_set_person_created_at();

CREATE OR REPLACE FUNCTION fn_set_person_created_at()
RETURNS TRIGGER
AS
$$  
BEGIN
    NEW.created_at := NOW();
    RETURN NEW;
END;
$$
LANGUAGE plpgsql;