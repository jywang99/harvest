--
-- PostgreSQL database cluster dump
--

-- Started on 2024-07-06 16:25:51 UTC

SET default_transaction_read_only = off;

SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;

--
-- Roles
--

CREATE ROLE postgres;
ALTER ROLE postgres WITH SUPERUSER INHERIT CREATEROLE CREATEDB LOGIN REPLICATION BYPASSRLS PASSWORD 'md55305adaac499dbbc6865a44e4aa5d8b4';






--
-- Databases
--

--
-- Database "template1" dump
--

\connect template1

--
-- PostgreSQL database dump
--

-- Dumped from database version 12.19 (Debian 12.19-1.pgdg120+1)
-- Dumped by pg_dump version 12.19

-- Started on 2024-07-06 16:25:51 UTC

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

-- Completed on 2024-07-06 16:25:51 UTC

--
-- PostgreSQL database dump complete
--

--
-- Database "postgres" dump
--

\connect postgres

--
-- PostgreSQL database dump
--

-- Dumped from database version 12.19 (Debian 12.19-1.pgdg120+1)
-- Dumped by pg_dump version 12.19

-- Started on 2024-07-06 16:25:51 UTC

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
-- TOC entry 7 (class 2615 OID 16384)
-- Name: media; Type: SCHEMA; Schema: -; Owner: postgres
--

CREATE SCHEMA media;


ALTER SCHEMA media OWNER TO postgres;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 205 (class 1259 OID 16389)
-- Name: collection; Type: TABLE; Schema: media; Owner: postgres
--

CREATE TABLE media.collection (
    id integer NOT NULL,
    path text NOT NULL,
    disp_name text,
    parent integer NOT NULL
);


ALTER TABLE media.collection OWNER TO postgres;

--
-- TOC entry 203 (class 1259 OID 16385)
-- Name: collection_id_seq; Type: SEQUENCE; Schema: media; Owner: postgres
--

CREATE SEQUENCE media.collection_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE media.collection_id_seq OWNER TO postgres;

--
-- TOC entry 3005 (class 0 OID 0)
-- Dependencies: 203
-- Name: collection_id_seq; Type: SEQUENCE OWNED BY; Schema: media; Owner: postgres
--

ALTER SEQUENCE media.collection_id_seq OWNED BY media.collection.id;


--
-- TOC entry 204 (class 1259 OID 16387)
-- Name: collection_parent_seq; Type: SEQUENCE; Schema: media; Owner: postgres
--

CREATE SEQUENCE media.collection_parent_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE media.collection_parent_seq OWNER TO postgres;

--
-- TOC entry 3006 (class 0 OID 0)
-- Dependencies: 204
-- Name: collection_parent_seq; Type: SEQUENCE OWNED BY; Schema: media; Owner: postgres
--

ALTER SEQUENCE media.collection_parent_seq OWNED BY media.collection.parent;


--
-- TOC entry 208 (class 1259 OID 16408)
-- Name: entry; Type: TABLE; Schema: media; Owner: postgres
--

CREATE TABLE media.entry (
    id integer NOT NULL,
    path text NOT NULL,
    disp_name text,
    "desc" text,
    is_dir boolean DEFAULT false NOT NULL,
    parent integer NOT NULL,
    thumb_paths text[]
);


ALTER TABLE media.entry OWNER TO postgres;

--
-- TOC entry 206 (class 1259 OID 16404)
-- Name: entry_id_seq; Type: SEQUENCE; Schema: media; Owner: postgres
--

CREATE SEQUENCE media.entry_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE media.entry_id_seq OWNER TO postgres;

--
-- TOC entry 3007 (class 0 OID 0)
-- Dependencies: 206
-- Name: entry_id_seq; Type: SEQUENCE OWNED BY; Schema: media; Owner: postgres
--

ALTER SEQUENCE media.entry_id_seq OWNED BY media.entry.id;


--
-- TOC entry 207 (class 1259 OID 16406)
-- Name: entry_parent_seq; Type: SEQUENCE; Schema: media; Owner: postgres
--

CREATE SEQUENCE media.entry_parent_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE media.entry_parent_seq OWNER TO postgres;

--
-- TOC entry 3008 (class 0 OID 0)
-- Dependencies: 207
-- Name: entry_parent_seq; Type: SEQUENCE OWNED BY; Schema: media; Owner: postgres
--

ALTER SEQUENCE media.entry_parent_seq OWNED BY media.entry.parent;


--
-- TOC entry 2863 (class 2604 OID 16392)
-- Name: collection id; Type: DEFAULT; Schema: media; Owner: postgres
--

ALTER TABLE ONLY media.collection ALTER COLUMN id SET DEFAULT nextval('media.collection_id_seq'::regclass);


--
-- TOC entry 2864 (class 2604 OID 16393)
-- Name: collection parent; Type: DEFAULT; Schema: media; Owner: postgres
--

ALTER TABLE ONLY media.collection ALTER COLUMN parent SET DEFAULT nextval('media.collection_parent_seq'::regclass);


--
-- TOC entry 2865 (class 2604 OID 16411)
-- Name: entry id; Type: DEFAULT; Schema: media; Owner: postgres
--

ALTER TABLE ONLY media.entry ALTER COLUMN id SET DEFAULT nextval('media.entry_id_seq'::regclass);


--
-- TOC entry 2867 (class 2604 OID 16413)
-- Name: entry parent; Type: DEFAULT; Schema: media; Owner: postgres
--

ALTER TABLE ONLY media.entry ALTER COLUMN parent SET DEFAULT nextval('media.entry_parent_seq'::regclass);


--
-- TOC entry 2869 (class 2606 OID 16398)
-- Name: collection pk_collection; Type: CONSTRAINT; Schema: media; Owner: postgres
--

ALTER TABLE ONLY media.collection
    ADD CONSTRAINT pk_collection PRIMARY KEY (id);


--
-- TOC entry 2871 (class 2606 OID 16418)
-- Name: entry pk_entry; Type: CONSTRAINT; Schema: media; Owner: postgres
--

ALTER TABLE ONLY media.entry
    ADD CONSTRAINT pk_entry PRIMARY KEY (id);


--
-- TOC entry 2872 (class 2606 OID 16399)
-- Name: collection fk_collection_self; Type: FK CONSTRAINT; Schema: media; Owner: postgres
--

ALTER TABLE ONLY media.collection
    ADD CONSTRAINT fk_collection_self FOREIGN KEY (parent) REFERENCES media.collection(id) NOT VALID;


--
-- TOC entry 2873 (class 2606 OID 16419)
-- Name: entry fk_entry_parent; Type: FK CONSTRAINT; Schema: media; Owner: postgres
--

ALTER TABLE ONLY media.entry
    ADD CONSTRAINT fk_entry_parent FOREIGN KEY (parent) REFERENCES media.collection(id);


-- Completed on 2024-07-06 16:25:51 UTC

--
-- PostgreSQL database dump complete
--

-- Completed on 2024-07-06 16:25:51 UTC

--
-- PostgreSQL database cluster dump complete
--

