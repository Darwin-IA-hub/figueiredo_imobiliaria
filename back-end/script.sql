/* FigueiredoLógico: */

CREATE TABLE Cliente (
    telefone VARCHAR PRIMARY KEY,
    nome VARCHAR,
    dataNascimento DATE,
    rendaBrutaMensal NUMERIC,
    quantidadeFilhos NUMERIC,
    anosCarteiraAssinada NUMERIC,
    teveSubsidio BOOLEAN,
    vaiUsarFGTS BOOLEAN
);

CREATE TABLE Financiamento (
    id NUMERIC PRIMARY KEY,
    tipo VARCHAR,
    fk_Cliente_telefone VARCHAR
);

CREATE TABLE Foto (
    id NUMERIC PRIMARY KEY,
    linkFoto VARCHAR,
    fk_Cliente_telefone VARCHAR
);

CREATE TABLE Imovel (
    id NUMERIC PRIMARY KEY,
    tipo VARCHAR,
    linkIPTU VARCHAR,
    cidade VARCHAR
);

CREATE TABLE Conjuge (
    id NUMERIC PRIMARY KEY,
    rendaBrutaConjuge NUMERIC,
    dataNascimentoConjuge DATE,
    fk_Cliente_telefone VARCHAR
);

CREATE TABLE Interesse (
    id NUMERIC PRIMARY KEY,
    interesseAtual VARCHAR,
    cidadeInteresse VARCHAR,
    intervaloPreco VARCHAR,
    observacao VARCHAR,
    tipoImovelInteresse VARCHAR,
    fk_Cliente_telefone VARCHAR,
    fk_Imovel_id NUMERIC,
    fk_Lancamento_id NUMERIC
);

CREATE TABLE ImovelVenda (
    fk_Imovel_id NUMERIC PRIMARY KEY,
    financiadoQuitado VARCHAR,
    docEmDia BOOLEAN,
    estaHabitado BOOLEAN
);

CREATE TABLE Lancamento (
    id NUMERIC PRIMARY KEY,
    cidade VARCHAR,
    nome VARCHAR
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
    FOREIGN KEY (fk_Imovel_id)
    REFERENCES Imovel (id)
    ON DELETE CASCADE;
 
ALTER TABLE Interesse ADD CONSTRAINT FK_Interesse_4
    FOREIGN KEY (fk_Lancamento_id)
    REFERENCES Lancamento (id)
    ON DELETE CASCADE;
 
ALTER TABLE ImovelVenda ADD CONSTRAINT FK_ImovelVenda_2
    FOREIGN KEY (fk_Imovel_id)
    REFERENCES Imovel (id)
    ON DELETE CASCADE;