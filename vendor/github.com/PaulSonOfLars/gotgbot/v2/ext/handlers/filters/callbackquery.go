package filters

import (
	"strings"

	"github.com/PaulSonOfLars/gotgbot/v2"
)

type CallbackQuery func(cq *gotgbot.CallbackQuery) bool

func (f CallbackQuery) And(f2 CallbackQuery) CallbackQuery {
	return func(cq *gotgbot.CallbackQuery) bool {
		return f(cq) && f2(cq)
	}
}

func (f CallbackQuery) Or(f2 CallbackQuery) CallbackQuery {
	return func(cq *gotgbot.CallbackQuery) bool {
		return f(cq) || f2(cq)
	}
}

func (f CallbackQuery) Not() CallbackQuery {
	return func(cq *gotgbot.CallbackQuery) bool {
		return !f(cq)
	}
}

func Prefix(prefix string) CallbackQuery {
	return func(cq *gotgbot.CallbackQuery) bool {
		return strings.HasPrefix(cq.Data, prefix)
	}
}

func CallbackUserID(id int64) CallbackQuery {
	return func(cq *gotgbot.CallbackQuery) bool {
		return cq.From.Id == id
	}
}

func GameName(name string) CallbackQuery {
	return func(cq *gotgbot.CallbackQuery) bool {
		return cq.GameShortName == name
	}
}
