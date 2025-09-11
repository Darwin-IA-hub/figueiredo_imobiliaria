/* Lógico_1: */

CREATE TABLE Cliente (
    telefone TEXT PRIMARY KEY,
    nomeCliente TEXT,
    dataNascimentoCliente DATE,
    rendaBrutaCliente NUMERIC,
    quantidadeFilhos NUMERIC,
    anosCarteiraAssinada NUMERIC,
    teveSubsidio BOOLEAN,
    vaiUsarFGTS BOOLEAN,
    possuiFinanciamento BOOLEAN
);

CREATE TABLE Financiamento (
    idFinanciamento SERIAL PRIMARY KEY,
    descricaoFinanciamento TEXT,
    fk_Cliente_telefone TEXT
);

CREATE TABLE Foto (
    idFoto SERIAL PRIMARY KEY,
    linkFoto TEXT,
    fk_Cliente_telefone TEXT
);

CREATE TABLE Imovel (
    idImovel SERIAL PRIMARY KEY,
    tipoImovel TEXT,
    linkIPTU TEXT,
    cidadeImovel TEXT
);

CREATE TABLE Conjuge (
    idConjuge SERIAL PRIMARY KEY,
    rendaBrutaMensalConjuge NUMERIC,
    dataNascimentoConjuge DATE,
    fk_Cliente_telefone TEXT
);

CREATE TABLE Interesse (
    idInteresse SERIAL PRIMARY KEY,
    interesseAtual TEXT,
    cidadeInteresse TEXT,
    intervaloPreco TEXT,
    observacao TEXT,
    tipoImovelInteresse TEXT,
    fk_Cliente_telefone TEXT,
    fk_Imovel_idImovel SERIAL,
    fk_Lancamento_idLancamento SERIAL
);

CREATE TABLE ImovelVenda (
    fk_Imovel_idImovel SERIAL PRIMARY KEY,
    financiadoQuitado TEXT,
    docEmDia BOOLEAN,
    estaHabitado BOOLEAN
);

CREATE TABLE Lancamento (
    idLancamento SERIAL PRIMARY KEY,
    cidadeLancamento TEXT,
    nomeLancamento TEXT
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