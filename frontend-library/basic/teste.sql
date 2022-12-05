--
-- PostgreSQL database dump
--

-- Dumped from database version 10.22 (Ubuntu 10.22-0ubuntu0.18.04.1)
-- Dumped by pg_dump version 10.22 (Ubuntu 10.22-0ubuntu0.18.04.1)

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

--
-- Name: raw; Type: SCHEMA; Schema: -; Owner: postgres
--

CREATE SCHEMA raw;


ALTER SCHEMA raw OWNER TO postgres;

--
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: Books; Type: TABLE; Schema: raw; Owner: postgres
--

CREATE TABLE raw.Books (
    id integer,
    content character varying(100),
    title character varying(100),
    author character varying(50),
    topic character varying(100)
);


ALTER TABLE raw.Books OWNER TO postgres;

--
-- Data for Name: Books; Type: TABLE DATA; Schema: raw; Owner: postgres
--

COPY raw.Books (id, content, title, author, topic) FROM stdin;
1	book	Harry Potter	JK Howling	Fiction and Magic
2	book	Mitologia Nórdica	Neil Gaiman	Mythology
3	book	God of War	Matthew Stover e Robert E. Vardeman	Mythology
\.


--
-- PostgreSQL database dump complete
--

