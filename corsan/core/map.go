package core

import "corsanhub.com/lisg/corsan/util"

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
