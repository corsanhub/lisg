package core

import "corsanhub.com/lisg/corsan/util"

func (reader *Reader) readForm() (MalType, error) {
	token, _ := reader.peek()
	if *reader.str == "\\" {
		return MalObject{v: "\\\\\\\\\\"}, nil
	} else {
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
		case "^":
			list, err := reader.readSpecialList(token)
			return list, err
		case "'":
			list, err := reader.readSpecialList(token)
			return list, err
		case "~":
			list, err := reader.readSpecialList(token)
			return list, err
		case "`":
			list, err := reader.readSpecialList(token)
			return list, err
		case "@":
			list, err := reader.readSpecialList(token)
			return list, err
		case "~@":
			list, err := reader.readSpecialList(token)
			return list, err

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

}
