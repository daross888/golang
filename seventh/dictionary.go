package seventh

var (
	UnknownWordError      = DictionaryError("could not find the word you are looking for")
	ErrorWordExist        = DictionaryError("word already exist in the dictionary")
	ErrorWordDoesNotExist = DictionaryError("could not perform operation because the word does not exist")
	ErrorNotFound         = DictionaryError("cannot delete word because it does not exist")
)

type DictionaryError string

func (e DictionaryError) Error() string {
	return string(e)
}

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", UnknownWordError
	}

	return definition, nil
}

func (d Dictionary) Add(definition, word string) error {
	_, err := d.Search(definition)

	switch err {
	case UnknownWordError:
		d[definition] = word
	case nil:
		return ErrorWordExist
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(definition, word string) error {
	_, err := d.Search(definition)

	switch err {
	case UnknownWordError:
		return ErrorWordDoesNotExist
	case nil:
		d[definition] = word
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)

	switch err {
	case UnknownWordError:
		return ErrorNotFound
	case nil:
		delete(d, word)
	default:
		return err
	}

	return nil
}

func Search(dictionary map[string]string, word string) string {
	return dictionary[word]
}
