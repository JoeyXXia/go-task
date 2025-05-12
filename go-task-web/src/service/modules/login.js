import xxRequest from "../request";

export function accountLogin(account) {
  return xxRequest.post({
    url: "/login",
    data: account,
  });
}

export function getUserInfoById(id) {
  return xxRequest.get({
    url: `/users/${id}`,
  });
}

export function getUserMenuByRoleInd(number) {
  return xxRequest.get({
    url: `/role/${id}/menu`,
  });
}
