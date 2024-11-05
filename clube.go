package main

import "errors"

var (
	ErrNumeroInvalidoJogadores = errors.New("o numero minimo de jogadores é 22")
	ErrTecnicoInvalido         = errors.New("tecnico invalido")
)

type Clube struct {
	nome    string
	estadio Estadio
	time    Time
	moral   int
	caixa   int
	divisao string
}

type Time struct {
	tecnico   Tecnico
	jogadores []Jogador
	// TODO: usar struct? ex: {defesa int, meioCampo int, ataque int} len = 10
	tatica string //343
}

func NewClube(nome string, estadio Estadio, time Time) Clube {
	return Clube{
		nome:    nome,
		estadio: estadio,
		time:    time,
	}
}

func NewTime(tecnico Tecnico, jogadores []Jogador) (Time, error) {
	const minimoJogadores = 11

	if len(jogadores) < minimoJogadores {
		return Time{}, ErrNumeroInvalidoJogadores
	}

	// TODO: validar perfil tecnico e nome de cada jogador

	if tecnico.nome == "" {
		return Time{}, ErrTecnicoInvalido
	}

	return Time{
		tecnico:   tecnico,
		jogadores: jogadores,
	}, nil
}
