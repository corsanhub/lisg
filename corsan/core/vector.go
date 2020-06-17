package core

import "corsanhub.com/lisg/corsan/util"

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
