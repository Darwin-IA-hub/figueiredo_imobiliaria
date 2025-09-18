/* Lógico_1: */

CREATE TABLE Cliente (
    telefone TEXT PRIMARY KEY,
    nomeCliente TEXT DEFAULT '' NOT NULL,
    dataNascimentoCliente DATE DEFAULT CURRENT_DATE,
    rendaBrutaCliente NUMERIC DEFAULT 0 NOT NULL,
    quantidadeFilhos NUMERIC DEFAULT 0 NOT NULL,
    anosCarteiraAssinada NUMERIC DEFAULT 0 NOT NULL,
    teveSubsidio BOOLEAN DEFAULT false NOT NULL,
    vaiUsarFGTS BOOLEAN DEFAULT false NOT NULL,
    possuiFinanciamento BOOLEAN DEFAULT false NOT NULL
);

CREATE TABLE Financiamento (
    idFinanciamento SERIAL PRIMARY KEY,
    descricaoFinanciamento TEXT DEFAULT '' NOT NULL,
    fk_Cliente_telefone TEXT
);

CREATE TABLE Foto (
    idFoto SERIAL PRIMARY KEY,
    linkFoto TEXT DEFAULT '' NOT NULL,
    fk_Cliente_telefone TEXT
);

CREATE TABLE Imovel (
    idImovel SERIAL PRIMARY KEY,
    tipoImovel TEXT DEFAULT '' NOT NULL,
    linkIPTU TEXT DEFAULT '' NOT NULL,
    cidadeImovel TEXT DEFAULT '' NOT NULL
);

CREATE TABLE Conjuge (
    idConjuge SERIAL PRIMARY KEY,
    rendaBrutaMensalConjuge NUMERIC DEFAULT 0 NOT NULL,
    dataNascimentoConjuge DATE DEFAULT CURRENT_DATE,
    fk_Cliente_telefone TEXT
);

CREATE TABLE Interesse (
    idInteresse SERIAL PRIMARY KEY,
    interesseAtual TEXT DEFAULT '' NOT NULL,
    cidadeInteresse TEXT DEFAULT '' NOT NULL,
    intervaloPreco TEXT  DEFAULT '' NOT NULL,
    observacao TEXT DEFAULT '' NOT NULL,
    tipoImovelInteresse TEXT DEFAULT '' NOT NULL,
    fk_Cliente_telefone TEXT ,
    fk_Imovel_idImovel INTEGER,
    fk_Lancamento_idLancamento INTEGER
);

CREATE TABLE ImovelVenda (
    fk_Imovel_idImovel INTEGER PRIMARY KEY,
    financiadoQuitado TEXT DEFAULT '' NOT NULL,
    docEmDia BOOLEAN DEFAULT false NOT NULL,
    estaHabitado BOOLEAN DEFAULT false NOT NULL
);

CREATE TABLE Lancamento (
    idLancamento SERIAL PRIMARY KEY,
    cidadeLancamento TEXT DEFAULT '' NOT NULL,
    nomeLancamento TEXT DEFAULT '' NOT NULL
);
 
ALTER TABLE Financiamento ADD CONSTRAINT FK_Financiamento_2
    FOREIGN KEY (fk_Cliente_telefone)
    REFERENCES Cliente (telefone)
    ON DELETE CASCADE;
 
ALTER TABLE Foto ADD CONSTRAINT FK_Foto_2
    FOREIGN KEY (fk_Cliente_telefone)
    REFERENCES Cliente (telefone)
    ON DELETE CASCADE;
 
ALTER TABLE Conjuge ADD CONSTRAINT FK_Conjuge_2
    FOREIGN KEY (fk_Cliente_telefone)
    REFERENCES Cliente (telefone)
    ON DELETE CASCADE;
 
ALTER TABLE Interesse ADD CONSTRAINT FK_Interesse_2
    FOREIGN KEY (fk_Cliente_telefone)
    REFERENCES Cliente (telefone)
    ON DELETE RESTRICT;
 
ALTER TABLE Interesse ADD CONSTRAINT FK_Interesse_3
    FOREIGN KEY (fk_Imovel_idImovel)
    REFERENCES Imovel (idImovel)
    ON DELETE CASCADE;
 
ALTER TABLE Interesse ADD CONSTRAINT FK_Interesse_4
    FOREIGN KEY (fk_Lancamento_idLancamento)
    REFERENCES Lancamento (idLancamento)
    ON DELETE CASCADE;
 
ALTER TABLE ImovelVenda ADD CONSTRAINT FK_ImovelVenda_2
    FOREIGN KEY (fk_Imovel_idImovel)
    REFERENCES Imovel (idImovel)
    ON DELETE CASCADE;


-- reset BD
TRUNCATE TABLE cliente RESTART IDENTITY CASCADE;
TRUNCATE TABLE imovel RESTART IDENTITY CASCADE;