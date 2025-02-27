-- migrations/000001_create_children_table.up.sql
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE groups (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    nome VARCHAR(255) NOT NULL,
    idade_minima INTEGER NOT NULL,
    idade_maxima INTEGER NOT NULL,
    descricao TEXT,
    criado_em TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    atualizado_em TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE TABLE needs (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    tipo VARCHAR(100) NOT NULL,
    descricao TEXT,
    criado_em TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    atualizado_em TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE TABLE allergies (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    tipo VARCHAR(100) NOT NULL,
    descricao TEXT,
    gravidade VARCHAR(50),
    criado_em TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    atualizado_em TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE TABLE children (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    nome VARCHAR(255) NOT NULL,
    data_nascimento DATE NOT NULL,
    sexo VARCHAR(20) NOT NULL,
    foto_url TEXT,
    grupo_id UUID REFERENCES groups(id),
    criado_em TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    atualizado_em TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE TABLE children_needs (
    crianca_id UUID REFERENCES children(id) ON DELETE CASCADE,
    necessidade_id UUID REFERENCES needs(id) ON DELETE CASCADE,
    PRIMARY KEY (crianca_id, necessidade_id)
);

CREATE TABLE children_allergies (
    crianca_id UUID REFERENCES children(id) ON DELETE CASCADE,
    alergia_id UUID REFERENCES allergies(id) ON DELETE CASCADE,
    PRIMARY KEY (crianca_id, alergia_id)
);

CREATE TABLE caretakers (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    nome VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE,
    telefone VARCHAR(50),
    tipo_telefone VARCHAR(20),
    auth0_id VARCHAR(255) UNIQUE,
    criado_em TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    atualizado_em TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE TABLE children_caretakers (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    crianca_id UUID REFERENCES children(id) ON DELETE CASCADE,
    responsavel_id UUID REFERENCES caretakers(id) ON DELETE CASCADE,
    tipo_relacao VARCHAR(50) NOT NULL,
    pode_retirar BOOLEAN DEFAULT TRUE,
    criado_em TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    atualizado_em TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    UNIQUE(crianca_id, responsavel_id)
);

CREATE TABLE volunteers (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    nome VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE,
    telefone VARCHAR(50),
    data_nascimento DATE,
    foto_url TEXT,
    verificacao_background BOOLEAN DEFAULT FALSE,
    data_verificacao DATE,
    auth0_id VARCHAR(255) UNIQUE,
    criado_em TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    atualizado_em TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Índices para melhorar a performance
CREATE INDEX idx_children_nome ON children(nome);
CREATE INDEX idx_children_data_nascimento ON children(data_nascimento);
CREATE INDEX idx_caretakers_nome ON caretakers(nome);
CREATE INDEX idx_caretakers_email ON caretakers(email);
CREATE INDEX idx_volunteers_nome ON volunteers(nome);
CREATE INDEX idx_volunteers_email ON volunteers(email);