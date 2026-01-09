package components

import (
	"strconv"

	twmerge "github.com/Oudwins/tailwind-merge-go"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

type InputFieldProps struct {
	Value    string
	Name     string
	Type     string
	Required bool
	Label    string
}

type TextareaProps struct {
	Value    string
	Name     string
	Required bool
	Rows     int
	Label    string
	Class    string
}

func InputField(props *InputFieldProps) Node {
	return Label(
		Class("grid gap-2"),
		Span(
			Class("font-bold"),
			Text(props.Label),
			If(
				props.Required,
				Span(Class("text-red-700 dark:text-red-300"), Text(" *")),
			),
		),
		Input(
			If(props.Type != "", Type(props.Type)),
			If(props.Required, Required()),
			Name(props.Name),
			Class("h-12 border-border border-solid border bg-surface rounded-sm px-3 text-lg outline-0"),
		),
	)
}

func TextareaField(props *TextareaProps) Node {
	return Label(
		Class("grid gap-2"),
		Span(
			Class("font-bold"),
			Text(props.Label),
			If(
				props.Required,
				Span(Class("text-red-700 dark:text-red-300"), Text(" *")),
			),
		),
		Textarea(
			Name(props.Name),
			Rows(strconv.Itoa(props.Rows)),
			Class(twmerge.Merge("bg-surface border-border border-solid border rounded-sm p-3 outline-0 text-lg", props.Class)),
		),
	)
}
