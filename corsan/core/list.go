package core

import "corsanhub.com/lisg/corsan/util"

const BASIC_LIST = "b"
const SPECIAL_LIST = "s"

func (reader *Reader) readList(id *string) (MalType, error) {
	list := MalList{id: id, t: BASIC_LIST}
	log.Debug(util.Xs("##-------------   new list  id: %v, %#v", id, "list"))

	for {
		token, err := reader.next()

		if token == nil || *token == "" || err != nil {
			if err != nil {
				//log.Warn(err.Error())
			}
			log.Debug("Breaking 📰 list ...")
			break
		}

		switch *token {
		case ")":
			reader.counter--
			log.Debug(util.Xs("##-------------   ret list  id: %v, %#v", id, "list"))
			return list, nil
		default:
			form, _ := reader.readForm()
			list.v = append(list.v, form)
			log.Debug(util.Xs("##-------------   appn list  id: %v, %#v, %#v", id, PrintStr(list), list.t))
			//waitForEnterKey("list-form :")
		}
	}

	// if list.t == SPECIAL_LIST {
	// 	return list, nil
	// }

	fn := util.TraceStr(0)
	return nil, NewError(fn, "unbalanced")
}

func (reader *Reader) readSpecialList(id *string) (MalType, error) {
	list := []MalType{}
	token, _ := reader.peek()

	switch *token {
	case "^":
		metaSymbol := MalSymbol{v: "with-meta"}
		list = append(list, metaSymbol)
	case "'":
		metaSymbol := MalSymbol{v: "quote"}
		list = append(list, metaSymbol)
	case "~":
		metaSymbol := MalSymbol{v: "unquote"}
		list = append(list, metaSymbol)
	case "`":
		metaSymbol := MalSymbol{v: "quasiquote"}
		list = append(list, metaSymbol)
	case "@":
		metaSymbol := MalSymbol{v: "deref"}
		list = append(list, metaSymbol)
	case "~@":
		metaSymbol := MalSymbol{v: "splice-unquote"}
		list = append(list, metaSymbol)
	default:
		metaSymbol := MalSymbol{v: "undefined"}
		list = append(list, metaSymbol)
	}

	log.Debug(util.Xs("##-------------   new spcl  id: %v, %#v", id, "spcl"))

	for {
		token, err := reader.next()

		if token == nil || *token == "" || err != nil {
			if err != nil {
				//log.Warn(err.Error())
			}
			log.Debug("Breaking 📰 spcl ...")
			break
		}

		form, _ := reader.readForm()
		if form != nil {
			log.Debug(util.Xs("##-------------   appn spcl  id: %v, %#v", id, "spcl"))
			list = append(list, form)
		}
	}

	specialList := MalList{id: id, t: SPECIAL_LIST}
	if *reader.tokens[0] == "^" {
		specialList.v = append(specialList.v, list[0])
		specialList.v = append(specialList.v, list[2])
		specialList.v = append(specialList.v, list[1])
	} else {
		specialList.v = list
	}

	log.Debug(util.Xs("##-------------    ret spcl  id: %v, %#v", id, PrintStr(specialList)))
	return specialList, nil
}
