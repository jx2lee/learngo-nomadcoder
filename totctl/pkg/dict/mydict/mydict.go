package mydict

import "errors"

// Dictionary type
type Dictionary map[string]string

var (
	errNotFound     = errors.New("Not Found..")
	errExitsAlready = errors.New("That words Already Exists..")
	errCantUpdate   = errors.New("That word Not Found..")
)

// Search for word
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}

// Add word
func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)
	switch err {
	case errNotFound:
		d[word] = def
	case nil:
		return errExitsAlready
	}
	return nil
}

// Update word
func (d Dictionary) Update(word, def string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		d[word] = def
	case errNotFound:
		return errCantUpdate
	}
	return nil
}

// Delete a word
func (d Dictionary) Delete(word string) {
	delete(d, word)
}
