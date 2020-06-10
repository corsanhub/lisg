package core

import (
	"fmt"

	"corsanhub.com/lisg/corsan/logging"
	"corsanhub.com/lisg/corsan/util"
)

var log = logging.Logger{Name: "core.reader"}

func NewError(fnName, str string) *MalError {
	return &MalError{f: fnName, e: str}
}

type Reader struct {
	position int
	counter  int
	tokens   []*string
}

func (reader *Reader) readList(id *string) (MalType, error) {
	list := MalList{id: id}
	log.Debug(util.Xs("##-------------   new list  id: %v, %#v", id, "list"))

	for {
		token, err := reader.next()

		if token == nil || *token == "" || err != nil {
			if err != nil {
				//log.Warn(err.Error())
			}
			log.Debug("Breaking 📰 ...")
			break
		}

		switch *token {
		case ")":
			reader.counter--
			//waitForEnterKeywaitForEnterKey("list-retr :")
			log.Debug(util.Xs("##-------------   ret list  id: %v, %#v", id, "list"))
			return list, nil
		default:
			form, _ := reader.readForm()
			list.v = append(list.v, form)
			//waitForEnterKey("list-form :")
		}
	}

	fn := util.TraceStr(0)
	return nil, NewError(fn, "unbalanced")
}

func (reader *Reader) readForm() (MalType, error) {
	token, _ := reader.peek()
	switch *token {
	case "(":
		reader.counter++
		list, err := reader.readList(token)
		//log.Debug(util.Xs("##-------------Return list : %#v", *&list))
		return list, err
	case "[":
		reader.counter++
		vector, err := reader.readVector(token)
		//log.Debug(util.Xs("##-------------Return list : %#v", *&list))
		return vector, err
	case "{":
		reader.counter++
		rmap, err := reader.readMap(token)
		//log.Debug(util.Xs("##-------------Return list : %#v", *&list))
		return rmap, err
	case ")":
		reader.counter--
		log.Debug("It's the end of a list.")
		return nil, nil
	case "]":
		reader.counter--
		log.Debug("It's the end of a vector.")
		return nil, nil
	case "}":
		reader.counter--
		log.Debug("It's the end of a map.")
		return nil, nil
	default:
		atom, err := reader.readAtom()
		log.Debug("It's an atom.")
		if atom != nil {
			return atom, err
		} else {
			fn := util.TraceStr(0)
			return nil, NewError(fn, "Form is not consistent")
		}
	}
}

func (reader *Reader) readVector(id *string) (MalType, error) {
	vector := MalVector{id: id}
	log.Debug(util.Xs("##----------   new vect  id: %v, %#v", id, "list"))

	for {
		token, err := reader.next()

		if token == nil || *token == "" || err != nil {
			if err != nil {
				//log.Warn(err.Error())
			}
			log.Debug("Breaking 📰 ...")
			break
		}

		switch *token {
		case "]":
			reader.counter--
			//waitForEnterKeywaitForEnterKey("list-retr :")
			log.Debug(util.Xs("##----------   ret vect  id: %v, %#v", id, "list"))
			return vector, nil
		default:
			form, _ := reader.readForm()
			vector.v = append(vector.v, form)
			//waitForEnterKey("list-form :")
		}
	}

	fn := util.TraceStr(0)
	return nil, NewError(fn, "Malformed list")
}

func (reader *Reader) readMap(id *string) (MalType, error) {
	rmap := MalMap{id: id}
	log.Debug(util.Xs("##--------------   new map  id: %v, %#v", id, "list"))

	for {
		token, err := reader.next()

		if token == nil || *token == "" || err != nil {
			if err != nil {
				//log.Warn(err.Error())
			}
			log.Debug("Breaking 📰 ...")
			break
		}

		switch *token {
		case "}":
			reader.counter--
			//waitForEnterKeywaitForEnterKey("list-retr :")
			log.Debug(util.Xs("##--------------   ret map  id: %v, %#v", id, "list"))
			return rmap, nil
		default:
			form, _ := reader.readForm()
			rmap.v = append(rmap.v, form)
			//waitForEnterKey("list-form :")
		}
	}

	fn := util.TraceStr(0)
	return nil, NewError(fn, "Malformed vector")
}

func CreateReader(str string) *Reader {
	tokens := Tokenize(str)
	log.Debug(util.Xs("tokens: %s\n", util.PointersToString(tokens, "'")))
	reader := &Reader{tokens: tokens}
	return reader
}

func ReadStr(str string) MalType {
	if &str != nil {
		reader := CreateReader(str)
		form, _ := reader.readForm()

		if reader.counter != 0 {
			return MalObject{v: "unbalanced"}
		}
		return form
	} else {
		return nil
	}
}

func TestXReader() {
	println("Testing reader ...")
	println("------------------------------------------------------------------------------------------------------------------------------")

	//str = "(a b)"
	str := "(()())"
	reader := CreateReader(str)
	form, _ := reader.readForm()
	fmt.Printf("form: %#v\n", form)

	println("------------------------------------------------------------------------------------------------------------------------------")

}
