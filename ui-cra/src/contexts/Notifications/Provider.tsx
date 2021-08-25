import React, { FC, useCallback, useEffect, useState } from "react";
import { useHistory } from "react-router-dom";
import { NotificationDialog } from "../../components/Layout/Notification";
import { Notification, NotificationData } from "./index";

const NotificationProvider: FC = ({ children }) => {
  const [notifications, setNotifications] = useState<NotificationData[] | []>(
    []
  );
  const [open, setOpen] = useState<boolean>(true);
  const history = useHistory();

  const clearNotifications = useCallback(() => {
    setNotifications([]);
    setOpen(true);
  }, []);

  useEffect(() => {
    setOpen(true);
    return history.listen(clearNotifications);
  }, [history, notifications, clearNotifications]);

  return (
    <Notification.Provider value={{ notifications, setNotifications }}>
      {open && notifications.length !== 0 && (
        <NotificationDialog onClose={setOpen} />
      )}
      {children}
    </Notification.Provider>
  );
};

export default NotificationProvider;
