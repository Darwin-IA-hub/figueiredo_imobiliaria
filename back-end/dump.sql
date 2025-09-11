--
-- PostgreSQL database dump
--

-- Dumped from database version 16.9 (Ubuntu 16.9-0ubuntu0.24.04.1)
-- Dumped by pg_dump version 16.9 (Ubuntu 16.9-0ubuntu0.24.04.1)

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
-- Name: fn_delete_old_msgs(); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.fn_delete_old_msgs() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    /*
      Remove mensagens do mesmo telefone
      registradas há mais de 2 minutos.
    */
    DELETE FROM Mensagens
      WHERE telefone = NEW.telefone
        AND data     < now() - interval '2 minutes';

    RETURN NEW;   -- mantém a linha recém-inserida
END;
$$;


ALTER FUNCTION public.fn_delete_old_msgs() OWNER TO postgres;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: cliente; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.cliente (
    telefone character varying NOT NULL,
    nomecliente character varying,
    dataNascimentoCliente DATE,
    rendaBrutaCliente NUMERIC,
    quantidadeFilhos NUMERIC,
    anosCarteiraAssinada NUMERIC,
    teveSubsidio BOOLEAN,
    vaiUsarFGTS BOOLEAN,
    possuiFinanciamento BOOLEAN
);


ALTER TABLE public.cliente OWNER TO postgres;

--
-- Name: clientesbloqueados; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.clientesbloqueados (
    idcliente text NOT NULL
);


ALTER TABLE public.clientesbloqueados OWNER TO postgres;

--
-- Name: contatos; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.contatos (
    telefone character varying NOT NULL,
    nome character varying,
    conversation_id character varying,
    ativo boolean
);


ALTER TABLE public.contatos OWNER TO postgres;

--
-- Name: conversas; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.conversas (
    telefone character varying(64),
    data character varying(10)
);


ALTER TABLE public.conversas OWNER TO postgres;

--
-- Name: mensagens; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.mensagens (
    id bigint NOT NULL,
    telefone character varying(64),
    conteudo character varying(1000),
    data timestamp with time zone DEFAULT now()
);


ALTER TABLE public.mensagens OWNER TO postgres;

--
-- Name: mensagens_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.mensagens_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.mensagens_id_seq OWNER TO postgres;

--
-- Name: mensagens_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.mensagens_id_seq OWNED BY public.mensagens.id;


--
-- Name: usuarios; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.usuarios (
    id bigint NOT NULL,
    unidade character varying(50),
    celular character varying(13) NOT NULL,
    senha character varying(16) NOT NULL,
    endereco character varying(100),
    ativo boolean DEFAULT true,
    role character varying(10) DEFAULT 'user'::character varying
);


ALTER TABLE public.usuarios OWNER TO postgres;

--
-- Name: usuarios_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.usuarios_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.usuarios_id_seq OWNER TO postgres;

--
-- Name: usuarios_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.usuarios_id_seq OWNED BY public.usuarios.id;


--
-- Name: mensagens id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mensagens ALTER COLUMN id SET DEFAULT nextval('public.mensagens_id_seq'::regclass);


--
-- Name: usuarios id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.usuarios ALTER COLUMN id SET DEFAULT nextval('public.usuarios_id_seq'::regclass);


--
-- Data for Name: cliente; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.cliente (telefone, nomecliente) FROM stdin;
\.


--
-- Data for Name: clientesbloqueados; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.clientesbloqueados (idcliente) FROM stdin;
\.


--
-- Data for Name: contatos; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.contatos (telefone, nome, conversation_id, ativo) FROM stdin;
\.


--
-- Data for Name: conversas; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.conversas (telefone, data) FROM stdin;
\.


--
-- Data for Name: mensagens; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.mensagens (id, telefone, conteudo, data) FROM stdin;
\.


--
-- Data for Name: usuarios; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.usuarios (id, unidade, celular, senha, endereco, ativo, role) FROM stdin;
\.


--
-- Name: mensagens_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.mensagens_id_seq', 2251, true);


--
-- Name: usuarios_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.usuarios_id_seq', 32, true);


--
-- Name: cliente cliente_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.cliente
    ADD CONSTRAINT cliente_pkey PRIMARY KEY (telefone);


--
-- Name: clientesbloqueados clientesbloqueados_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.clientesbloqueados
    ADD CONSTRAINT clientesbloqueados_pkey PRIMARY KEY (idcliente);


--
-- Name: contatos contatos_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.contatos
    ADD CONSTRAINT contatos_pkey PRIMARY KEY (telefone);


--
-- Name: mensagens mensagens_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mensagens
    ADD CONSTRAINT mensagens_pkey PRIMARY KEY (id);


--
-- Name: usuarios usuarios_celular_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.usuarios
    ADD CONSTRAINT usuarios_celular_key UNIQUE (celular);


--
-- Name: usuarios usuarios_endereco_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.usuarios
    ADD CONSTRAINT usuarios_endereco_key UNIQUE (endereco);


--
-- Name: usuarios usuarios_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.usuarios
    ADD CONSTRAINT usuarios_pkey PRIMARY KEY (id);


--
-- Name: idx_mensagens_telefone_data; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_mensagens_telefone_data ON public.mensagens USING btree (telefone, data);


--
-- Name: mensagens tg_delete_old_msgs; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER tg_delete_old_msgs AFTER INSERT ON public.mensagens FOR EACH ROW EXECUTE FUNCTION public.fn_delete_old_msgs();


--
-- Name: cliente cliente_telefone_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.cliente
    ADD CONSTRAINT cliente_telefone_fkey FOREIGN KEY (telefone) REFERENCES public.contatos(telefone) ON DELETE CASCADE;


--
-- Name: cliente fk_cliente_contato; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.cliente
    ADD CONSTRAINT fk_cliente_contato FOREIGN KEY (telefone) REFERENCES public.contatos(telefone) ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--