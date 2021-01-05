package app

import (
	"bufio"
	"log"
	"os"
)

type history struct {
	items    []string
	filename string
	maxLines int
}

func NewHistory(filename string, maxLines int) (*history, error) {
	items := []string{}

	file, err := os.OpenFile(filename, os.O_CREATE, 0755)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		items = append(items, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return &history{
		items:    items,
		filename: filename,
		maxLines: maxLines,
	}, nil
}

func (h *history) Insert(s string) {
	h.items = append(h.items, s)

	if len(h.items) > h.maxLines {
		h.items = h.items[0:h.maxLines]
	}

	// for now, flush after each insert
	err := h.Flush()
	if err != nil {
		log.Printf("unable to write to history file: %s", err.Error())
	}
}

func (h *history) Items() []string {
	return h.items
}

func (h *history) Flush() error {
	f, err := os.Create(h.filename)
	if err != nil {
		return err
	}
	defer f.Close()

	for _, v := range h.items {
		_, err = f.WriteString(v + "\n")
		if err != nil {
			return err
		}
	}

	return nil
}
