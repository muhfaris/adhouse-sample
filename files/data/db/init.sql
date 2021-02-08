-- CREATE TABLE "users" ----------------------------------------
CREATE TABLE "public"."users" (
	"id" Integer NOT NULL,
	"username" Character Varying( 2044 ) NOT NULL,
	"password" Character Varying( 2044 ) NOT NULL,
	"first_name" Character Varying( 2044 ) NOT NULL,
	"last_name" Character Varying( 2044 ) NOT NULL,
	CONSTRAINT "unique_users_id" UNIQUE( "id" ) );
 ;
-- -------------------------------------------------------------

INSERT INTO public.users VALUES (0, 'admin01', 'admin01', 'adminadhouse', '01');


-- CREATE TABLE "product" --------------------------------------
CREATE TABLE "public"."products" (
	"id" Integer NOT NULL,
	"name" Character Varying( 2044 ) NOT NULL,
	"qty" Integer NOT NULL,
	CONSTRAINT "unique_products_id" UNIQUE( "id" ) );
 ;
-- -------------------------------------------------------------

INSERT INTO public.products VALUES (0, 'Laptop 01', 50);
INSERT INTO public.products VALUES (1, 'Laptop 02', 150);
INSERT INTO public.products VALUES (2, 'Laptop 03', 10);
INSERT INTO public.products VALUES (3, 'Laptop 04', 140);
