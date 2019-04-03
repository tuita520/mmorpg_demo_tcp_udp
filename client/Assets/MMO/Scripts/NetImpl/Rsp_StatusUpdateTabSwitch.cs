﻿using System;
using msg;

namespace ghbc.Net
{
	public class Rsp_StatusUpdateTabSwitch:S2C_Handler
	{
		public override void execute (short cmd, short ErrCode, byte[] payloads)
		{
			DebugTool.LogPurple ("角色切换到射击状态");
			base.execute (cmd, ErrCode, payloads);
			var data = getData<Rspn_LeaveOffOthers> ();
			var playerList = data.Player;
			facade.DispatchEvent (new CEvent (Event_Player.LeaveOffOthers.ToString (), new object[]{ playerList }));
		}
	}
}

