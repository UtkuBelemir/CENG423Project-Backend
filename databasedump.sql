--
-- PostgreSQL database dump
--

-- Dumped from database version 9.6.11
-- Dumped by pg_dump version 11.1

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: aybu_student_system; Type: DATABASE; Schema: -; Owner: aybu
--

CREATE DATABASE aybu_student_system WITH TEMPLATE = template0 ENCODING = 'UTF8' LC_COLLATE = 'tr_TR.utf8' LC_CTYPE = 'tr_TR.utf8';


ALTER DATABASE aybu_student_system OWNER TO aybu;

\connect aybu_student_system

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: public; Type: SCHEMA; Schema: -; Owner: postgres
--

CREATE SCHEMA public;


ALTER SCHEMA public OWNER TO postgres;

--
-- Name: SCHEMA public; Type: COMMENT; Schema: -; Owner: postgres
--

COMMENT ON SCHEMA public IS 'standard public schema';


SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: advertisements; Type: TABLE; Schema: public; Owner: aybu
--

CREATE TABLE public.advertisements (
    record_id character varying(40) DEFAULT public.uuid_generate_v1() NOT NULL,
    title character varying(350) NOT NULL,
    description text,
    price double precision,
    category character varying(50),
    image bytea,
    owner character varying(40),
    type character varying(40),
    status integer DEFAULT 0
);


ALTER TABLE public.advertisements OWNER TO aybu;

--
-- Name: users; Type: TABLE; Schema: public; Owner: aybu
--

CREATE TABLE public.users (
    username character varying(40) NOT NULL,
    password character varying(255) NOT NULL,
    last_name character varying(40) NOT NULL,
    email character varying(350) NOT NULL,
    first_name character varying(255) NOT NULL,
    gender integer NOT NULL,
    role character varying(40)
);


ALTER TABLE public.users OWNER TO aybu;

--
-- Data for Name: advertisements; Type: TABLE DATA; Schema: public; Owner: aybu
--



--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: aybu
--

INSERT INTO public.users (username, password, last_name, email, first_name, gender, role) VALUES ('admin', '$2a$10$2DfygDzORTAdyq67owa4D.6j/.qOemAsUK8m8xgEamSQaaDTMRDum', 'Admin', 'admin@aybu.com', 'System', 0, 'admin');


--
-- Name: advertisements advertisements_pk; Type: CONSTRAINT; Schema: public; Owner: aybu
--

ALTER TABLE ONLY public.advertisements
    ADD CONSTRAINT advertisements_pk PRIMARY KEY (record_id);


--
-- Name: users users_pk; Type: CONSTRAINT; Schema: public; Owner: aybu
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pk PRIMARY KEY (username);


--
-- Name: advertisements_record_id_uindex; Type: INDEX; Schema: public; Owner: aybu
--

CREATE UNIQUE INDEX advertisements_record_id_uindex ON public.advertisements USING btree (record_id);


--
-- Name: users_username_uindex; Type: INDEX; Schema: public; Owner: aybu
--

CREATE UNIQUE INDEX users_username_uindex ON public.users USING btree (username);


--
-- PostgreSQL database dump complete
--

