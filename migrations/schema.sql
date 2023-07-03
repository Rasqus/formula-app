--
-- PostgreSQL database dump
--

-- Dumped from database version 15.3 (Debian 15.3-1.pgdg110+1)
-- Dumped by pg_dump version 15.0

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

SET default_table_access_method = heap;

--
-- Name: DRIVER; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."DRIVER" (
    "ID" integer NOT NULL,
    "NAME" character varying(255) NOT NULL,
    "SURNAME" character varying(255) NOT NULL,
    "COUNTRY" character varying(255) NOT NULL,
    "START_NUMBER" integer NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    "TEAM" character varying(255) NOT NULL
);


ALTER TABLE public."DRIVER" OWNER TO postgres;

--
-- Name: DRIVER_ID_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public."DRIVER_ID_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."DRIVER_ID_seq" OWNER TO postgres;

--
-- Name: DRIVER_ID_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public."DRIVER_ID_seq" OWNED BY public."DRIVER"."ID";


--
-- Name: RACE; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."RACE" (
    "ID" integer NOT NULL,
    "NAME" character varying(255) NOT NULL,
    "DETAILS" character varying(300),
    "SEASON_ID" integer NOT NULL,
    "TRACK_ID" integer NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public."RACE" OWNER TO postgres;

--
-- Name: GRAND_PRIX_ID_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public."GRAND_PRIX_ID_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."GRAND_PRIX_ID_seq" OWNER TO postgres;

--
-- Name: GRAND_PRIX_ID_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public."GRAND_PRIX_ID_seq" OWNED BY public."RACE"."ID";


--
-- Name: PLACE; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."PLACE" (
    "ID" integer NOT NULL,
    "RACE_ID" integer NOT NULL,
    "DRIVER_ID" integer NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    "PLACE" integer NOT NULL
);


ALTER TABLE public."PLACE" OWNER TO postgres;

--
-- Name: GRAND_PRIX_TO_DRIVER_ID_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public."GRAND_PRIX_TO_DRIVER_ID_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."GRAND_PRIX_TO_DRIVER_ID_seq" OWNER TO postgres;

--
-- Name: GRAND_PRIX_TO_DRIVER_ID_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public."GRAND_PRIX_TO_DRIVER_ID_seq" OWNED BY public."PLACE"."ID";


--
-- Name: SEASON; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."SEASON" (
    "ID" integer NOT NULL,
    "YEAR" integer NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public."SEASON" OWNER TO postgres;

--
-- Name: SEASON_ID_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public."SEASON_ID_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."SEASON_ID_seq" OWNER TO postgres;

--
-- Name: SEASON_ID_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public."SEASON_ID_seq" OWNED BY public."SEASON"."ID";


--
-- Name: TRACK; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."TRACK" (
    "ID" integer NOT NULL,
    "NAME" character varying(255) NOT NULL,
    "DETAILS" character varying(300) NOT NULL,
    "COUNTRY" character varying(255) NOT NULL,
    "CITY" character varying(255) NOT NULL,
    "LENGTH" integer NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public."TRACK" OWNER TO postgres;

--
-- Name: TRACK_ID_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public."TRACK_ID_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."TRACK_ID_seq" OWNER TO postgres;

--
-- Name: TRACK_ID_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public."TRACK_ID_seq" OWNED BY public."TRACK"."ID";


--
-- Name: schema_migration; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.schema_migration (
    version character varying(14) NOT NULL
);


ALTER TABLE public.schema_migration OWNER TO postgres;

--
-- Name: DRIVER ID; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."DRIVER" ALTER COLUMN "ID" SET DEFAULT nextval('public."DRIVER_ID_seq"'::regclass);


--
-- Name: PLACE ID; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."PLACE" ALTER COLUMN "ID" SET DEFAULT nextval('public."GRAND_PRIX_TO_DRIVER_ID_seq"'::regclass);


--
-- Name: RACE ID; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."RACE" ALTER COLUMN "ID" SET DEFAULT nextval('public."GRAND_PRIX_ID_seq"'::regclass);


--
-- Name: SEASON ID; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."SEASON" ALTER COLUMN "ID" SET DEFAULT nextval('public."SEASON_ID_seq"'::regclass);


--
-- Name: TRACK ID; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."TRACK" ALTER COLUMN "ID" SET DEFAULT nextval('public."TRACK_ID_seq"'::regclass);


--
-- Name: DRIVER DRIVER_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."DRIVER"
    ADD CONSTRAINT "DRIVER_pkey" PRIMARY KEY ("ID");


--
-- Name: PLACE GRAND_PRIX_TO_DRIVER_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."PLACE"
    ADD CONSTRAINT "GRAND_PRIX_TO_DRIVER_pkey" PRIMARY KEY ("ID");


--
-- Name: RACE GRAND_PRIX_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."RACE"
    ADD CONSTRAINT "GRAND_PRIX_pkey" PRIMARY KEY ("ID");


--
-- Name: SEASON SEASON_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."SEASON"
    ADD CONSTRAINT "SEASON_pkey" PRIMARY KEY ("ID");


--
-- Name: TRACK TRACK_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."TRACK"
    ADD CONSTRAINT "TRACK_pkey" PRIMARY KEY ("ID");


--
-- Name: schema_migration schema_migration_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.schema_migration
    ADD CONSTRAINT schema_migration_pkey PRIMARY KEY (version);


--
-- Name: schema_migration_version_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX schema_migration_version_idx ON public.schema_migration USING btree (version);


--
-- Name: RACE GRAND_PRIX_SEASON_ID_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."RACE"
    ADD CONSTRAINT "GRAND_PRIX_SEASON_ID_fkey" FOREIGN KEY ("SEASON_ID") REFERENCES public."SEASON"("ID") ON DELETE CASCADE;


--
-- Name: PLACE GRAND_PRIX_TO_DRIVER_DRIVER_ID_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."PLACE"
    ADD CONSTRAINT "GRAND_PRIX_TO_DRIVER_DRIVER_ID_fkey" FOREIGN KEY ("DRIVER_ID") REFERENCES public."DRIVER"("ID") ON DELETE CASCADE;


--
-- Name: PLACE GRAND_PRIX_TO_DRIVER_GRANDPRIX_ID_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."PLACE"
    ADD CONSTRAINT "GRAND_PRIX_TO_DRIVER_GRANDPRIX_ID_fkey" FOREIGN KEY ("RACE_ID") REFERENCES public."RACE"("ID") ON DELETE CASCADE;


--
-- Name: RACE GRAND_PRIX_TRACK_ID_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."RACE"
    ADD CONSTRAINT "GRAND_PRIX_TRACK_ID_fkey" FOREIGN KEY ("TRACK_ID") REFERENCES public."TRACK"("ID") ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

