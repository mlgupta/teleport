/**
 * Teleport
 * Copyright (C) 2024  Gravitational, Inc.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

import { useState, useMemo, useEffect, useCallback } from 'react';
import { isBefore, isAfter, formatDistanceToNowStrict } from 'date-fns';
import styled from 'styled-components';
import { Alert, Box, Flex, Indicator, Text } from 'design';

import { Notification as NotificationIcon, BellRinging } from 'design/Icon';
import Logger from 'shared/libs/logger';
import { useRefClickOutside } from 'shared/hooks/useRefClickOutside';
import { HoverTooltip } from 'shared/components/ToolTip';

import {
  useInfiniteScroll,
  useKeyBasedPagination,
} from 'shared/hooks/useInfiniteScroll';
import { IGNORE_CLICK_CLASSNAME } from 'shared/hooks/useRefClickOutside/useRefClickOutside';

import { useStore } from 'shared/libs/stores';

import { useTeleport } from 'teleport';
import useStickyClusterId from 'teleport/useStickyClusterId';
import { Dropdown } from 'teleport/components/Dropdown';

import { ButtonIconContainer } from 'teleport/TopBar/Shared';

import {
  LocalNotificationGroupedKind,
  LocalNotificationKind,
  Notification as NotificationType,
} from 'teleport/services/notifications';

import {
  Notification as AccessListNotification,
  LocalNotificationStates,
} from 'teleport/stores/storeNotifications';

import { Notification } from './Notification';

const PAGE_SIZE = 15;

const logger = Logger.create('Notifications');

const NOTIFICATION_DROPDOWN_ID = 'tb-notifications-dropdown';

export function Notifications({ iconSize = 24 }: { iconSize?: number }) {
  const ctx = useTeleport();
  const { clusterId } = useStickyClusterId();
  const store = useStore(ctx.storeNotifications);
  const [userLastSeenNotification, setUserLastSeenNotification] =
    useState<Date>();

  const [combinedNotifications, setCombinedNotifications] = useState<
    NotificationType[]
  >([]);

  // Access list review reminder notifications aren't currently supported by the native notifications
  // system, and so they won't be returned by the prior request. They are instead generated by the frontend
  // and stored in a local  store on the frontend. We retrieve these separately and then combine them with
  // the "real" notifications from the backend.
  const localNotifications = useMemo(
    () =>
      accessListStoreNotificationsToNotifications(
        store.getNotifications(),
        store.getNotificationStates()
      ),
    [store.state]
  );

  const {
    resources: notifications,
    fetch,
    attempt,
  } = useKeyBasedPagination({
    fetchMoreSize: PAGE_SIZE,
    initialFetchSize: PAGE_SIZE,
    fetchFunc: useCallback(
      async paginationParams => {
        const response = await ctx.notificationService.fetchNotifications({
          clusterId,
          startKey: paginationParams.startKey,
          limit: paginationParams.limit,
        });

        setUserLastSeenNotification(response.userLastSeenNotification);

        return {
          agents: response.notifications,
          startKey: response.nextKey,
        };
      },
      [clusterId, ctx.notificationService]
    ),
  });

  // Fetch first page on first render.
  useEffect(() => {
    fetch();
  }, []);

  useEffect(() => {
    setCombinedNotifications([...localNotifications, ...notifications]);
  }, [localNotifications, notifications]);

  const { setTrigger } = useInfiniteScroll({
    fetch,
  });

  const [view, setView] = useState<View>('All');
  const [open, setOpen] = useState(false);

  const ref = useRefClickOutside<HTMLDivElement>({ open, setOpen });

  function onIconClick() {
    if (!open) {
      setOpen(true);

      if (localNotifications.length) {
        store.markNotificationsAsSeen(
          localNotifications.map(notif => notif.id)
        );
      }

      if (notifications.length) {
        const latestNotificationTime = notifications[0].createdDate;
        // If the current userLastSeenNotification is already set to the most recent notification's time, don't do anything.
        if (userLastSeenNotification === latestNotificationTime) {
          return;
        }

        const previousLastSeenTime = userLastSeenNotification;

        // Update the visual state right away for a snappier UX.
        setUserLastSeenNotification(latestNotificationTime);

        ctx.notificationService
          .upsertLastSeenNotificationTime(clusterId, {
            time: latestNotificationTime,
          })
          .then(res => setUserLastSeenNotification(res.time))
          .catch(err => {
            setUserLastSeenNotification(previousLastSeenTime);
            logger.error(`Notification last seen time update failed.`, err);
          });
      }
    } else {
      setOpen(false);
    }
  }

  const unseenNotifsCount = combinedNotifications.filter(notif => {
    if (notif.localNotification) {
      const seenNotifications = store.getNotificationStates().seen;

      return !seenNotifications.includes(notif.id);
    }

    return isBefore(userLastSeenNotification, notif.createdDate);
  }).length;

  function removeNotification(notificationId: string) {
    const notificationsCopy = [...combinedNotifications];
    const index = notificationsCopy.findIndex(
      notif => notif.id == notificationId
    );
    notificationsCopy.splice(index, 1);

    setCombinedNotifications(notificationsCopy);
  }

  function markNotificationAsClicked(notificationId: string) {
    const newNotifications = combinedNotifications.map(notification => {
      return notification.id === notificationId
        ? { ...notification, clicked: true }
        : notification;
    });

    setCombinedNotifications(newNotifications);
  }

  return (
    <NotificationButtonContainer
      ref={ref}
      data-testid="tb-notifications"
      className={IGNORE_CLICK_CLASSNAME}
    >
      <HoverTooltip
        anchorOrigin={{ vertical: 'bottom', horizontal: 'center' }}
        transformOrigin={{ vertical: 'top', horizontal: 'center' }}
        tipContent="Notifications"
        css={`
          height: 100%;
        `}
      >
        <ButtonIconContainer
          onClick={onIconClick}
          onKeyUp={e => (e.key === 'Enter' || e.key === ' ') && onIconClick()}
          data-testid="tb-notifications-button"
          open={open}
          role="button"
          tabIndex={0}
          aria-label="Notifications"
          aria-haspopup="menu"
          aria-controls={NOTIFICATION_DROPDOWN_ID}
          aria-expanded={open}
        >
          {unseenNotifsCount > 0 && (
            <UnseenBadge data-testid="tb-notifications-badge">
              {unseenNotifsCount >= 9 ? '9+' : unseenNotifsCount}
            </UnseenBadge>
          )}
          <NotificationIcon
            color={open ? 'text.main' : 'text.muted'}
            size={iconSize}
          />
        </ButtonIconContainer>
      </HoverTooltip>

      <NotificationsDropdown
        open={open}
        id={NOTIFICATION_DROPDOWN_ID}
        data-testid={NOTIFICATION_DROPDOWN_ID}
        role="menu"
      >
        <Header view={view} setView={setView} />
        {attempt.status === 'failed' && (
          <Box px={3}>
            <Alert>Could not load notifications: {attempt.statusText}</Alert>
          </Box>
        )}
        {attempt.status === 'success' && combinedNotifications.length === 0 && (
          <EmptyState />
        )}
        <NotificationsList>
          <>
            {!!combinedNotifications.length &&
              combinedNotifications.map(notif => (
                <Notification
                  notification={notif}
                  key={notif.id}
                  view={view}
                  closeNotificationsList={() => setOpen(false)}
                  markNotificationAsClicked={markNotificationAsClicked}
                  removeNotification={removeNotification}
                />
              ))}
            {open && <div ref={setTrigger} />}
            {attempt.status === 'processing' && (
              <Flex
                width="100%"
                justifyContent="center"
                alignItems="center"
                mt={2}
              >
                <Indicator />
              </Flex>
            )}
          </>
        </NotificationsList>
      </NotificationsDropdown>
    </NotificationButtonContainer>
  );
}

function Header({
  view,
  setView,
}: {
  view: View;
  setView: (view: View) => void;
}) {
  return (
    <Box
      css={`
        padding: 0px ${p => p.theme.space[3]}px;
        width: 100%;
      `}
    >
      <Flex
        css={`
          flex-direction: column;
          box-sizing: border-box;
          gap: 12px;
          border-bottom: 1px solid
            ${p => p.theme.colors.interactive.tonal.neutral[2]};
          padding-bottom: ${p => p.theme.space[3]}px;
          margin-bottom: ${p => p.theme.space[3]}px;
        `}
      >
        <Text typography="dropdownTitle">Notifications</Text>
        <Flex gap={2}>
          <ViewButton selected={view === 'All'} onClick={() => setView('All')}>
            All
          </ViewButton>
          <ViewButton
            selected={view === 'Unread'}
            onClick={() => setView('Unread')}
          >
            Unread
          </ViewButton>
        </Flex>
      </Flex>
    </Box>
  );
}

function EmptyState() {
  return (
    <Flex
      flexDirection="column"
      alignItems="center"
      justifyContent="center"
      width="100%"
      height="100%"
      mt={4}
      mb={4}
    >
      <Flex
        css={`
          align-items: center;
          justify-content: center;
          height: 88px;
          width: 88px;
          background-color: ${p => p.theme.colors.interactive.tonal.neutral[0]};
          border-radius: ${p => p.theme.radii[7]}px;
          border: 1px solid ${p => p.theme.colors.interactive.tonal.neutral[1]};
        `}
      >
        <BellRinging size={40} />
      </Flex>
      <Text
        mt={4}
        css={`
          font-weight: 500;
          font-size: 18px;
          line-height: 24px;
          text-align: center;
        `}
      >
        You currently have no notifications.
      </Text>
    </Flex>
  );
}

/** accessListStoreNotificationsToNotifications converts a list of access list notifications from the notifications store into the primary
 * Notification type used by the notifications list.
 */
function accessListStoreNotificationsToNotifications(
  accessListNotifs: AccessListNotification[],
  notificationStates: LocalNotificationStates
): NotificationType[] {
  const today = new Date();

  /** dueNotifications are the notifications for access lists which are due for review in the future. */
  const dueNotifications = accessListNotifs
    .filter(notif => isAfter(notif.date, today))
    // Sort by earliest dates.
    .sort((a, b) => {
      return a.date.getTime() - b.date.getTime();
    });

  /** overdueNotifications are the notifications for access lists which are overdue for review. */
  const overdueNotifications = accessListNotifs
    .filter(notif => isBefore(notif.date, today))
    .sort((a, b) => {
      return a.date.getTime() - b.date.getTime();
    });

  const processedDueNotifications: NotificationType[] = [];
  const processedOverdueNotifications: NotificationType[] = [];

  // If there are 2 or less access list notifications due for review, then we will return a notification per access list.
  // If there are more than 2, then we return one "grouped" notification for all of them. This is to prevent clutter in the notifications list
  // in case there are many access lists due for review.
  if (dueNotifications.length <= 2) {
    // Process and add them to the final processed array.
    dueNotifications.forEach(notif => {
      const numDays = formatDistanceToNowStrict(notif.date);
      const titleText = `Access list '${notif.item.resourceName}' needs your review within ${numDays}.`;

      processedDueNotifications.push({
        localNotification: true,
        title: titleText,
        id: notif.id,
        subKind: LocalNotificationKind.AccessList,
        clicked: notif.clicked,
        createdDate: today,
        labels: [{ name: 'redirect-route', value: notif.item.route }],
      });
    });
  } else {
    const mostUrgentNotif = dueNotifications[0];
    const numDays = formatDistanceToNowStrict(mostUrgentNotif.date);
    const titleText = `${dueNotifications.length} of your access lists require review, the most urgent of which is due in ${numDays}.`;

    // The ID for this combined notification is <first notification in the list id>-<last notification in the list id>-<length of list>.
    const id = `${dueNotifications[0].id}-${dueNotifications[dueNotifications.length - 1].id}-${dueNotifications.length}`;

    const clicked = notificationStates.clicked.includes(id);

    processedDueNotifications.push({
      localNotification: true,
      title: titleText,
      id,
      subKind: LocalNotificationGroupedKind.AccessListGrouping,
      clicked,
      createdDate: today,
      labels: [],
    });
  }

  if (overdueNotifications.length <= 2) {
    // Process and add them to the final processed array.
    overdueNotifications.forEach(notif => {
      const numDays = formatDistanceToNowStrict(notif.date);
      const titleText = `Your review of access list '${notif.item.resourceName}' is overdue by ${numDays}.`;

      processedOverdueNotifications.push({
        localNotification: true,
        title: titleText,
        id: notif.id,
        subKind: LocalNotificationKind.AccessList,
        clicked: notif.clicked,
        createdDate: today,
        labels: [{ name: 'redirect-route', value: notif.item.route }],
      });
    });
  } else {
    const titleText = `${overdueNotifications.length} of your access lists are overdue for review.`;

    // The ID for this combined notification is <first notification in the list id>-<last notification in the list id>-<length of list>.
    const id = `${overdueNotifications[0].id}-${overdueNotifications[overdueNotifications.length - 1].id}-${overdueNotifications.length}`;

    const clicked = notificationStates.clicked.includes(id);

    processedOverdueNotifications.push({
      localNotification: true,
      title: titleText,
      id,
      subKind: LocalNotificationGroupedKind.AccessListGrouping,
      clicked,
      createdDate: today,
      labels: [],
    });
  }

  return [...processedOverdueNotifications, ...processedDueNotifications];
}

const NotificationsDropdown = styled(Dropdown)`
  width: 450px;
  padding: 0px;
  padding-top: ${p => p.theme.space[3]}px;
  align-items: center;
  height: 80vh;

  right: -40px;
  @media screen and (min-width: ${p => p.theme.breakpoints.small}px) {
    right: -52px;
  }
  @media screen and (min-width: ${p => p.theme.breakpoints.large}px) {
    right: -140px;
  }
`;

const ViewButton = styled.div<{ selected: boolean }>`
  cursor: pointer;
  align-items: center;
  // TODO(rudream): Clean up radii order in sharedStyles.
  border-radius: 36px;
  display: flex;
  width: fit-content;
  padding: ${p => p.theme.space[1]}px ${p => p.theme.space[3]}px;
  justify-content: space-around;
  font-size: 14px;
  font-weight: 300;
  color: ${props =>
    props.selected
      ? props.theme.colors.text.primaryInverse
      : props.theme.colors.text.muted};
  background-color: ${props =>
    props.selected ? props.theme.colors.brand : 'transparent'};

  .selected {
    color: ${props => props.theme.colors.text.primaryInverse};
    background-color: ${props => props.theme.colors.brand};
    transition: color 0.2s ease-in 0s;
  }
`;

export type View = 'All' | 'Unread';

const NotificationsList = styled.div`
  box-sizing: border-box;
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  gap: ${p => p.theme.space[2]}px;
  width: 100%;
  max-height: 100%;
  overflow-y: auto;
  padding: ${p => p.theme.space[3]}px;
  padding-top: 0px;
  // Subtract the width of the scrollbar from the right padding.
  padding-right: ${p => `${p.theme.space[3] - 8}px`};

  ::-webkit-scrollbar-thumb {
    background-color: ${p => p.theme.colors.interactive.tonal.neutral[2]};
    border-radius: ${p => p.theme.radii[2]}px;
    // Trick to make the scrollbar thumb 2px narrower than the track.
    border: 2px solid transparent;
    background-clip: padding-box;
  }

  ::-webkit-scrollbar {
    width: 8px;
    border-radius: ${p => p.theme.radii[2]}px;
    border-radius: ${p => p.theme.radii[2]}px;
    background-color: ${p => p.theme.colors.interactive.tonal.neutral[0]};
  }

  .notification {
    width: ${p => `${450 - p.theme.space[3] * 2}px`};
  }
`;

const NotificationButtonContainer = styled.div`
  position: relative;
  height: 100%;
`;

const UnseenBadge = styled.div`
  position: absolute;
  width: 16px;
  height: 16px;
  font-size: 10px;
  border-radius: 100%;
  color: ${p => p.theme.colors.text.primaryInverse};
  background-color: ${p => p.theme.colors.buttons.warning.default};
  margin-top: -21px;
  margin-right: -13px;
  display: flex;
  align-items: center;
  justify-content: center;
`;
