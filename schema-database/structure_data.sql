--
-- PostgreSQL database dump
--

-- Dumped from database version 11.2 (Debian 11.2-1.pgdg90+1)
-- Dumped by pg_dump version 13.1

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

--
-- Name: products; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.products (
    id integer NOT NULL,
    name character varying(100) NOT NULL,
    price numeric(12,0) NOT NULL,
    qty integer DEFAULT 0 NOT NULL
);


ALTER TABLE public.products OWNER TO admin;

--
-- Name: items_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.items_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.items_id_seq OWNER TO admin;

--
-- Name: items_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE public.items_id_seq OWNED BY public.products.id;


--
-- Name: orders; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.orders (
    id integer NOT NULL,
    order_trx character varying(20) NOT NULL,
    user_id integer NOT NULL,
    product_name character varying DEFAULT '100'::character varying NOT NULL,
    price numeric(10,0) NOT NULL,
    qty integer NOT NULL,
    total_price numeric(12,0) NOT NULL,
    created_at timestamp with time zone NOT NULL
);


ALTER TABLE public.orders OWNER TO admin;

--
-- Name: orders_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.orders_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.orders_id_seq OWNER TO admin;

--
-- Name: orders_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE public.orders_id_seq OWNED BY public.orders.id;


--
-- Name: users; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.users (
    id integer NOT NULL,
    name character varying(150) NOT NULL,
    email character varying(100) NOT NULL,
    address character varying NOT NULL
);


ALTER TABLE public.users OWNER TO admin;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.users_id_seq OWNER TO admin;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- Name: orders id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.orders ALTER COLUMN id SET DEFAULT nextval('public.orders_id_seq'::regclass);


--
-- Name: products id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.products ALTER COLUMN id SET DEFAULT nextval('public.items_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- Data for Name: orders; Type: TABLE DATA; Schema: public; Owner: admin
--



--
-- Data for Name: products; Type: TABLE DATA; Schema: public; Owner: admin
--

INSERT INTO public.products VALUES
	(1, 'laptop', 15000000, 5),
	(2, 'gamis', 100000, 20),
	(3, 'lipstik', 35000, 50),
	(4, 'selai coklat', 300000, 15);


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: admin
--

INSERT INTO public.users VALUES
	(1, 'Andy', 'andy.winarko@gmail.com', 'kemang'),
	(2, 'Budi', 'budi.oktoviyan@gmail.com', 'gambir'),
	(3, 'Catur', 'catur.teguh@gmail.com', 'bekasi timur'),
	(4, 'Dina', 'ramadina182@gmail.com', 'dago'),
	(5, 'Esti', 'estimasi@yahoo.com', 'priuk');


--
-- Name: items_id_seq; Type: SEQUENCE SET; Schema: public; Owner: admin
--

SELECT pg_catalog.setval('public.items_id_seq', 8, true);


--
-- Name: orders_id_seq; Type: SEQUENCE SET; Schema: public; Owner: admin
--

SELECT pg_catalog.setval('public.orders_id_seq', 1, false);


--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: admin
--

SELECT pg_catalog.setval('public.users_id_seq', 5, true);


--
-- Name: products unique_items_id; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.products
    ADD CONSTRAINT unique_items_id PRIMARY KEY (id);


--
-- Name: orders unique_orders_id; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.orders
    ADD CONSTRAINT unique_orders_id PRIMARY KEY (id);


--
-- Name: users unique_users_id; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT unique_users_id PRIMARY KEY (id);


--
-- Name: index_email; Type: INDEX; Schema: public; Owner: admin
--

CREATE INDEX index_email ON public.users USING btree (email);


--
-- Name: index_name; Type: INDEX; Schema: public; Owner: admin
--

CREATE INDEX index_name ON public.products USING btree (name);


--
-- Name: index_name1; Type: INDEX; Schema: public; Owner: admin
--

CREATE INDEX index_name1 ON public.users USING btree (name);


--
-- Name: index_order_trx; Type: INDEX; Schema: public; Owner: admin
--

CREATE INDEX index_order_trx ON public.orders USING btree (order_trx);


--
-- Name: index_user_id; Type: INDEX; Schema: public; Owner: admin
--

CREATE INDEX index_user_id ON public.orders USING btree (user_id);


--
-- PostgreSQL database dump complete
--

